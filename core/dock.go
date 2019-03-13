package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
)

type Dock interface {
	GetId() string
	Watch() error
}

// SentinelDock knows when to launch new sentinels by
// watching a scheduling queue
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
func (d *SentinelDock) GetId() string {
	return d.id
}

// Watch observes a queue in order to know when launch new Sentinels
// TODO add tests
// TODO add logging
func (sd *SentinelDock) Watch() (err error) {
	fmt.Println("Watching... ")

	// Hardcoding a stock to test sentinel
	// TODO: replace this hardcoded schedule with something more flexible.
	var schedule = NewSchedule("PETR3.SA", "1min")
	LaunchSentinel(sd.ctx, schedule)

	return nil
}

// TODO remove mocked behaviour
// TODO add tests
// TODO improve logging
func LaunchSentinel(context Context, schedule *Schedule) {
	var logger = context.GetLogger()

	var sentinel, err = NewStockSentinel(context, schedule)
	if err != nil {
		logger.Error("Fail to instantiate sentinel",
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("method", "main"),
			zap.String("error", err.Error()))
		os.Exit(1)
	}

	// Creating AlphaVantage client instance
	alphaVantage := client.NewAlphaVantage(context.GetSentinelConfig().AlphaVantageKey)

	// Running sentinel
	executionId, err := sentinel.Run(alphaVantage)
	if err != nil {
		logger.Error("Fail to run sentinel",
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("method", "main"),
			zap.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Sentinel have run successfully",
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "main"))

	// Trying to kill sentinel
	if err := sentinel.Kill(); err != nil {
		logger.Error("Fail to kill sentinel",
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Sentinel have terminated successfully",
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "main"))
}
