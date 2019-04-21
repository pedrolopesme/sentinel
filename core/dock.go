package core

import (
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// Dock represents a common interface to build all specific dock implementations
type Dock interface {
	GetId() string
	Watch() error
}

// SentinelDock knows when to launch new sentinels by
// watching a event stream
type SentinelDock struct {
	id  string
	ctx Context
}

// NewSentinelDock provides a working SentinelDock to a given app context
// TODO add tests
func NewSentinelDock(ctx Context) *SentinelDock {
	var executionId = uuid.Must(uuid.NewV4()).String()

	return &SentinelDock{
		id:  executionId,
		ctx: ctx,
	}
}

// GetId returns dock id (or execution id, if you prefer)
func (sd *SentinelDock) GetId() string {
	return sd.id
}

// Watch observes a queue in order to know when launch new Sentinels
// TODO add tests
func (sd *SentinelDock) Watch() (err error) {
	var logger = sd.ctx.Logger()

	logger.Info("Watching stocks",
		zap.String("dockId", sd.GetId()),
		zap.String("method", "main"))

	// Hardcoding a stock to test sentinel
	// TODO: replace this hardcoded schedule with something more flexible.
	var schedule = NewSchedule("PETR3.SA", "1min")

	if err := LaunchSentinel(sd.id, sd.ctx, schedule); err != nil {
		logger.Error("Fail to lunch Sentinel",
			zap.String("dockId", sd.GetId()),
			zap.String("method", "Watch"),
			zap.String("error", err.Error()))
		return err
	}

	logger.Info("Put have run successfully",
		zap.String("dockId", sd.GetId()),
		zap.String("method", "main"))

	return nil
}

// LaunchSentinel creates and put a Sentinel to check stock price changes
// TODO remove mocked behaviour
// TODO add tests
func LaunchSentinel(dockId string, context Context, schedule *Schedule) (err error) {
	var logger = context.Logger()

	sentinel, err := NewStockSentinel(context, schedule)
	if err != nil {
		logger.Error("Fail to instantiate sentinel",
			zap.String("dockId", dockId),
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("method", "LaunchSentinel"),
			zap.String("error", err.Error()))
		return err
	}

	// Creating AlphaVantage client instance
	alphaVantage := client.NewAlphaVantage(context.SentinelConfig().AlphaVantageKey)

	// Running sentinel
	executionId, err := sentinel.Run(alphaVantage)
	if err != nil {
		logger.Error("Fail to run sentinel",
			zap.String("dockId", dockId),
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("method", "LaunchSentinel"),
			zap.String("error", err.Error()))
		return err
	}

	logger.Info("Sentinel have run successfully",
		zap.String("dockId", dockId),
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "LaunchSentinel"))

	// Trying to kill sentinel
	if err := sentinel.Kill(); err != nil {
		logger.Error("Fail to kill sentinel",
			zap.String("dockId", dockId),
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("method", "LaunchSentinel"),
			zap.String("error", err.Error()))
		return err
	}

	logger.Info("Sentinel have terminated successfully",
		zap.String("dockId", dockId),
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "LaunchSentinel"))

	return nil
}
