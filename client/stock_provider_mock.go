package client

import (
	"github.com/stretchr/testify/mock"
)

type MockedStockProvider struct {
	mock.Mock
}

func (sp *MockedStockProvider) GetName() string {
	args := sp.Called()
	return args.String(0)
}

func (sp *MockedStockProvider) GetStocks(stock string, timeFrame string) (*StocksByTime, error) {
	args := sp.Called()
	return args.Get(0).(*StocksByTime), args.Error(1)
}
