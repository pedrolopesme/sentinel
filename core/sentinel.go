package core

import (
	"fmt"
	"time"

	"github.com/nats-io/go-nats-streaming"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/pedrolopesme/sentinel/models"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

const (
	NATS_STOCKS_SUBJECT = "stocks"
)

type Sentinel interface {
	// Id returns an unique identifier to the sentinel
	Id() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run(stockProvider client.StockProvider) (string, error)
}

type StockSentinel struct {
	id       string
	ctx      Context
	schedule *Schedule
	natsConn stan.Conn
}

// Id returns a unique identifier to the sentinel
func (s *StockSentinel) Id() string {
	return s.id
}

// Id returns a unique identifier to the sentinel
func (s *StockSentinel) Run(stockProvider client.StockProvider) (string, error) {
	var (
		executionId = uuid.Must(uuid.NewV4()).String()
		logger      = s.ctx.Logger()
	)

	logger.Info("Running StockSentinel",
		zap.String("sentinelId", s.Id()),
		zap.String("executionId", executionId))

	stocks, err := stockProvider.GetStocks(s.schedule.Stock, s.schedule.TimeFrame)
	if err != nil {
		logger.Error("Cant get stocks",
			zap.String("sentinelId", s.Id()),
			zap.String("executionId", executionId),
			zap.String("provider", stockProvider.GetName()),
			zap.String("error", err.Error()))
		return "", err
	}

	if err := s.publishStocks(executionId, stockProvider, stocks); err != nil {
		logger.Error("Cant publish stocks",
			zap.String("sentinelId", s.Id()),
			zap.String("executionId", executionId),
			zap.String("provider", stockProvider.GetName()),
			zap.String("error", err.Error()))
		return "", err
	}

	return executionId, nil
}

// TODO extract to its own structure. Is "Stocks Publisher" a good name?
func (s *StockSentinel) publishStocks(executionId string, stockProvider client.StockProvider, stocks map[time.Time]models.StockTier) (err error) {
	var (
		logger          = s.ctx.Logger()
		stockNATSClient = s.ctx.StockNats().GetConnection()
	)

	logger.Info(fmt.Sprintf("Found %v stocks. Publishing them to stocks queue", len(stocks)),
		zap.String("sentinelId", s.Id()),
		zap.String("provider", stockProvider.GetName()),
		zap.String("executionId", executionId))

	// TODO add tests
	// TODO what if publish fails? What about a retry logic?
	// TODO format message properly
	// Why not goroutines instead of a linear publishing?
	for timeFrame, stock := range stocks {
		logger.Info("Publishing stock",
			zap.String("sentinelId", s.Id()),
			zap.String("stock", s.schedule.Stock),
			zap.String("timeFrame", timeFrame.String()),
			zap.String("executionId", executionId))

		payload := fmt.Sprint(timeFrame, ">>>", stock.Price.High)
		before := time.Now()
		if err = stockNATSClient.Publish(NATS_STOCKS_SUBJECT, []byte(payload)); err != nil {
			logger.Error("Impossible to publish stock",
				zap.String("sentinelId", s.Id()),
				zap.String("stock", s.schedule.Stock),
				zap.String("timeFrame", timeFrame.String()),
				zap.String("millisecondsSpend", time.Since(before).String()),
				zap.String("executionId", executionId),
				zap.String("error", err.Error()))
		} else {
			logger.Info("Stock published",
				zap.String("sentinelId", s.Id()),
				zap.String("stock", s.schedule.Stock),
				zap.String("timeFrame", timeFrame.String()),
				zap.String("millisecondsSpend", time.Since(before).String()),
				zap.String("executionId", executionId))
		}
	}

	return
}

// NewSentinel is a base Sentinel build
func NewStockSentinel(ctx Context, schedule *Schedule) (sentinel *StockSentinel, err error) {
	sentinel = &StockSentinel{
		id:       uuid.Must(uuid.NewV4()).String(),
		schedule: schedule,
		ctx:      ctx,
	}

	ctx.Logger().Info("Sentinel created", zap.String("sentinelId", sentinel.Id()))
	return
}
