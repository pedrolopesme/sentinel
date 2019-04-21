package core

import (
	"fmt"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/pedrolopesme/sentinel/models"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

const (
	NATS_STOCKS_SUBJECT = "stocks"
)

type Sentinel interface {
	// GetId returns an unique identifier to the sentinel
	GetId() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run(stockProvider client.StockProvider) (string, error)
}

type StockSentinel struct {
	id       string
	ctx      Context
	schedule *Schedule
	natsConn stan.Conn
}

// GetId returns a unique identifier to the sentinel
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
func (s *StockSentinel) Run(stockProvider client.StockProvider) (string, error) {
	var (
		executionId = uuid.Must(uuid.NewV4()).String()
		logger      = s.ctx.Logger()
	)

	logger.Info("Running StockSentinel",
		zap.String("sentinelId", s.GetId()),
		zap.String("executionId", executionId))

	stocks, err := stockProvider.GetStocks(s.schedule.Stock, s.schedule.TimeFrame)
	if err != nil {
		logger.Error("Cant get stocks",
			zap.String("sentinelId", s.GetId()),
			zap.String("executionId", executionId),
			zap.String("provider", stockProvider.GetName()),
			zap.String("error", err.Error()))
		return "", err
	}

	if err := s.publishStocks(executionId, stockProvider, stocks); err != nil {
		logger.Error("Cant publish stocks",
			zap.String("sentinelId", s.GetId()),
			zap.String("executionId", executionId),
			zap.String("provider", stockProvider.GetName()),
			zap.String("error", err.Error()))
		return "", err
	}

	return executionId, nil
}

func (s *StockSentinel) publishStocks(executionId string, stockProvider client.StockProvider, stocks map[time.Time]models.StockTier) (err error) {
	var (
		logger          = s.ctx.Logger()
		stockNATSClient = s.ctx.StockNats().GetConnection()
	)

	logger.Info(fmt.Sprintf("Found %v stocks. Publishing them to stocks queue", len(stocks)),
		zap.String("sentinelId", s.GetId()),
		zap.String("provider", stockProvider.GetName()),
		zap.String("executionId", executionId))

	// TODO add tests
	// TODO what if publish fails? What about a retry logic?
	// TODO format message properly
	for timeFrame, stock := range stocks {
		logger.Info("Publishing stock",
			zap.String("sentinelId", s.GetId()),
			zap.String("stock", s.schedule.Stock),
			zap.String("timeFrame", timeFrame.String()),
			zap.String("executionId", executionId))

		payload := fmt.Sprint(timeFrame, ">>>", stock.Price.High)
		before := time.Now()
		if err = stockNATSClient.Publish(NATS_STOCKS_SUBJECT, []byte(payload)); err != nil {
			logger.Error("Impossible to publish stock",
				zap.String("sentinelId", s.GetId()),
				zap.String("stock", s.schedule.Stock),
				zap.String("timeFrame", timeFrame.String()),
				zap.String("millisecondsSpend", time.Since(before).String()),
				zap.String("executionId", executionId),
				zap.String("error", err.Error()))
		} else {
			logger.Info("Stock published",
				zap.String("sentinelId", s.GetId()),
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

	ctx.Logger().Info("Sentinel created", zap.String("sentinelId", sentinel.GetId()))
	return
}
