package core

import (
	"github.com/pedrolopesme/sentinel/client"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockedContext struct {
	mock.Mock
}

func (mc MockedContext) SentinelConfig() *SentinelConfig {
	args := mc.Called()
	return args.Get(0).(*SentinelConfig)
}

func (mc MockedContext) StockNats() client.NATSServer {
	args := mc.Called()
	return args.Get(0).(client.NATSServer)
}

func (mc MockedContext) Logger() *zap.Logger {
	args := mc.Called()
	return  args.Get(0).(*zap.Logger)
}
