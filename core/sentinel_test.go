package core

import (
	"github.com/pedrolopesme/sentinel/client"
	"github.com/pkg/errors"
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

func TestSentinelShouldReturnItsId(t *testing.T) {
	setup(t)
	dummyId := "some-dummyId"
	sentinel := StockSentinel{id: dummyId}
	assert.Equal(dummyId, sentinel.Id())
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

func TestSentinelsShouldStopProperlyWhenAErrorOccurWhileGettingStocks(t *testing.T) {
	var (
		schedule = NewSchedule("foo", "bar")
		stocks   = make(client.StocksByTime)
	)

	setup(t)
	contextMock.On("Logger").Return(dummyLogger).Times(2)
	stockProviderMock.On("GetName").Return("boo").Once()
	stockProviderMock.On("GetStocks", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&stocks, errors.New("some error")).
		Once()

	sentinel, _ := NewStockSentinel(contextMock, schedule)
	_, err := sentinel.Run(stockProviderMock)
	assert.NotNil(err)
}

func TestSentinelsShouldStopProperlyWhenNoStocksWereFoundButNoErrorWereReturned(t *testing.T) {
	var (
		schedule = NewSchedule("foo", "bar")
		stocks   = make(client.StocksByTime)
	)

	setup(t)

	// TODO: my NatsServer should not expose Nat's internal connection. It should
	// keep nats connection internally and expose the main funs (eg Publish/Consume).
	natServerMock.On("GetConnection").Return(nil)
	contextMock.On("Logger").Return(dummyLogger)
	contextMock.On("StockNats").Return(natServerMock)
	stockProviderMock.On("GetName").Return("boo").Once()
	stockProviderMock.On("GetStocks", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&stocks, nil).
		Once()

	sentinel, _ := NewStockSentinel(contextMock, schedule)
	_, err := sentinel.Run(stockProviderMock)
	assert.NotNil(err)
}

func TestSentinelsShouldStopProperlyWhenItWasImpossibleToPublishStocks(t *testing.T) {
	setup(t)
	assert.True(false)
}

func TestSentinelsShouldReturnItsExecutionAndNoErrosWhenItHasRunSuccessfully(t *testing.T) {
	setup(t)
	assert.True(false)
}
