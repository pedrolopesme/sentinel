package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

const (
	NATS_STOCKS_SUBJECT = "stocks"
)

type Sentinel interface {
	// GetId returns a unique identifier to the sentinel
	GetId() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run(stockProvider client.StockProvider) (string, error)

	// Kill stops a Sentinel
	Kill() error
}

type StockSentinel struct {
	id       string
	ctx      Context
	schedule *Schedule
}

// GetId returns a unique identifier to the sentinel
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
func (s *StockSentinel) Run(stockProvider client.StockProvider) (string, error) {
	var (
		logger      = s.ctx.GetLogger()
		executionId = uuid.Must(uuid.NewV4()).String()
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
	}

	logger.Info(fmt.Sprintf("Found %v stocks. Publishing them to stocks queue", len(stocks)),
		zap.String("sentinelId", s.GetId()),
		zap.String("provider", stockProvider.GetName()),
		zap.String("executionId", executionId))

	logger.Info("Connecting to Stocks Queue",
		zap.String("sentinelId", s.GetId()),
		zap.String("executionId", executionId))
	beforeConnect := time.Now()
	var stockNATSClient = s.ctx.GetStockNats().GetConnection()
	logger.Info("Connected to Stocks Queue",
		zap.String("sentinelId", s.GetId()),
		zap.String("millisecondsSpent", time.Since(beforeConnect).String()),
		zap.String("executionId", executionId))

	defer func() {
		logger.Info("Disconnecting from Stocks Queue",
			zap.String("sentinelId", s.GetId()),
			zap.String("executionId", executionId))
		before := time.Now()
		if err := stockNATSClient.Close(); err != nil {
			logger.Error("Error to disconnect from Stocks Queue",
				zap.String("sentinelId", s.GetId()),
				zap.String("millisecondsSpent", time.Since(before).String()),
				zap.String("executionId", executionId),
				zap.String("error", err.Error()))
		} else {
			logger.Info("Disconnected from Stocks Queue",
				zap.String("sentinelId", s.GetId()),
				zap.String("millisecondsSpent", time.Since(before).String()),
				zap.String("executionId", executionId))
		}
	}()

	// TODO extract it somewhere else
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

	return executionId, nil
}

// Kill stops a Sentinel
// TODO add tests
func (s *StockSentinel) Kill() error {
	return nil
}

// NewSentinel
// TODO add tests
func NewStockSentinel(ctx Context, schedule *Schedule) (sentinel *StockSentinel, err error) {
	sentinel = &StockSentinel{
		id:       uuid.Must(uuid.NewV4()).String(),
		schedule: schedule,
		ctx:      ctx,
	}

	ctx.GetLogger().Info("Sentinel created", zap.String("sentinelId", sentinel.GetId()))
	return
}
