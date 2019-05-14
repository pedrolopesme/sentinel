package client

import (
	"github.com/pedrolopesme/sentinel/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockedStockProvider struct {
	mock.Mock
}

func (sp *MockedStockProvider) GetName() string {
	args := sp.Called()
	return args.String(0)
}

func (sp *MockedStockProvider) GetStocks(stock string, timeFrame string) (map[time.Time]models.StockTier, error) {
	args := sp.Called()
	return args.Get(0).(map[time.Time]models.StockTier), args.Error(1)
}
