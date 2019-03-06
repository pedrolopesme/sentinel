package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
	"os"
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
	id         string
	config     *SentinelConfig
	schedule   *Schedule
	stocksNats client.NATSServer
}

// GetId returns a unique identifier to the sentinel
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
// TODO add log
func (s *StockSentinel) Run(stockProvider client.StockProvider) (string, error) {
	var executionId = uuid.Must(uuid.NewV4()).String()
	fmt.Println("Running StockSentinel ", s.GetId(), " - execution ", executionId)

	stocks, err := stockProvider.GetStocks(s.schedule.Stock, s.schedule.TimeFrame)
	if err != nil {
		fmt.Println("Cant get stocks due to", err.Error())
	}
	fmt.Printf("Found %v stocks. Publishing those to stocks queue \n", len(stocks))

	fmt.Println("Connecting to Stocks Queue")
	var stockNATSClient = s.stocksNats.GetConnection()
	fmt.Println("Connected to Stocks Queue")

	defer func() {
		fmt.Println("Disconnecting from Stocks Queue")
		stockNATSClient.Close()
		fmt.Println("Disconnected from Stocks Queue")
	}()

	// TODO extract it somewhere else
	// TODO add tests
	// TODO what if publish fails? What about a retry logic?
	// TODO format message properly
	for k, y := range stocks {
		stock := fmt.Sprint(k, ">>>", y.Price.High)
		fmt.Printf("Publishing stock %v\n", k)
		if err = stockNATSClient.Publish(NATS_STOCKS_SUBJECT, []byte(stock)); err != nil {
			fmt.Println("Cant public stocks to queue due to", err.Error())
			os.Exit(1)
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
// TODO add logging
func NewStockSentinel(config *SentinelConfig, schedule *Schedule) (sentinel *StockSentinel, err error) {
	clientID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Cant get stocks due to", err.Error())
	}

	stockNATS, err := client.NewNATSServer(config.NATSStocksClusterID, clientID.String(), config.NATSStocksURI)
	if err != nil {
		return nil, err
	}

	return &StockSentinel{
		id:         uuid.Must(uuid.NewV4()).String(),
		schedule:   schedule,
		config:     config,
		stocksNats: stockNATS,
	}, nil
}
