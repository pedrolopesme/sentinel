package core

import (
	"github.com/pedrolopesme/sentinel/client"
	assert2 "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var (
	// Mocks
	contextMock       *MockedContext
	dockMock          *MockedDock
	scheduleMock      *MockedSchedule
	sentinelMock      *MockedSentinel
	stockProviderMock *client.MockedStockProvider

	// Dummy objects
	dummyLogger *zap.Logger

	// Testify objects
	assert *assert2.Assertions
)

func setup(t *testing.T) {
	dummyLogger = zap.NewNop()
	assert = assert2.New(t)

	// Mocks
	contextMock = &MockedContext{}
	stockProviderMock = &client.MockedStockProvider{}
}
