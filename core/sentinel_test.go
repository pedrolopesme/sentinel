package core

import (
	"github.com/pedrolopesme/sentinel/client"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedSentinel struct {
	mock.Mock
}

func (ms *MockedSentinel) Id() string {
	args := ms.Called()
	return args.String(0)
}

func (ms *MockedSentinel) Run(stockProvider client.StockProvider) (string, error) {
	args := ms.Called(stockProvider)
	return args.String(0), args.Error(1)
}

func TestNewSentinelShouldReturnASentinelWithAUniqueId(t *testing.T) {
	var (
		schedule = NewSchedule("foo", "bar")
	)

	setup(t)
	contextMock.On("Logger").Return(dummyLogger).Times(2)
	firstSentinel, _ := NewStockSentinel(contextMock, schedule)
	secondSentinel, _ := NewStockSentinel(contextMock, schedule)
	assert.NotNil(firstSentinel)
	assert.NotNil(secondSentinel)
	assert.NotEqual(firstSentinel.id, secondSentinel.id)
}

func TestSentinelShouldReturnItsId(t *testing.T) {
	setup(t)
	dummyId := "some-dummyId"
	sentinel := StockSentinel{id: dummyId}
	assert.Equal(dummyId, sentinel.Id())
}
