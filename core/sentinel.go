package core

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
)

type Sentinel interface {
	// GetId returns a unique identifier to the sentinel
	GetId() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run() (string, error)

	// Kill stops a Sentinel
	Kill() error
}

// TODO extract alphaVantageKey from here.
// TODO decouple sentinel logic from http consume logic.
// TODO Alphavantage deserves its own client.
type StockSentinel struct {
	id              string
	alphaVantageKey string
	schedule        *Schedule
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
// TODO add log
// TODO extract all AlphaVantage retrieval to its own client
func (s *StockSentinel) Run() (string, error) {
	var executionId = uuid.Must(uuid.NewV4()).String()
	fmt.Println("Running StockSentinel ", s.GetId(), " - execution ", executionId)

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%v&interval=%v&outputsize=full&apikey=%v", s.schedule.Stock, s.schedule.TimeFrame, s.alphaVantageKey)
	client := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(body))
	return executionId, nil
}

// Kill stops a Sentinel
// TODO add tests
func (s *StockSentinel) Kill() error {
	return nil
}

// NewSentinel
// TODO add tests
func NewStockSentinel(alphaVantage string, schedule *Schedule) (sentinel *StockSentinel) {
	return &StockSentinel{
		id:              uuid.Must(uuid.NewV4()).String(),
		schedule:        schedule,
		alphaVantageKey: alphaVantage,
	}
}
