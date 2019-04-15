package core

import (
	"github.com/pedrolopesme/sentinel/client"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedSentinel struct {
	mock.Mock
}

func (ms *MockedSentinel) GetId() string {
	args := ms.Called()
	return args.String(0)
}

func (ms *MockedSentinel) Run(stockProvider client.StockProvider) (string, error) {
	args := ms.Called(stockProvider)
	return args.String(0), args.Error(1)
}

func (ms *MockedSentinel) Kill() error {
	args := ms.Called()
	return args.Error(1)
}

func TestNewSentinelShouldReturnASentinelWithAUniqueId(t *testing.T) {
	var (
		assert   = assert2.New(t)
		schedule = NewSchedule("foo", "bar")
		config   = SentinelConfig{}
		ctx, _   = NewAppContext(&config)
	)

	firstSentinel, _ := NewStockSentinel(ctx, schedule)
	secondSentinel, _ := NewStockSentinel(ctx, schedule)
	assert.NotNil(firstSentinel)
	assert.NotNil(secondSentinel)
	assert.NotEqual(firstSentinel.id, secondSentinel.id)
}
