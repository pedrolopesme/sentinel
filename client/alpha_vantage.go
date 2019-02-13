package client

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/models"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	URL               = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%v&interval=%v&outputsize=full&apikey=%v"
	HTTP_TIMEOUT_SECS = 10
)

// AlphaVantage is a client to https://www.alphavantage.co/
type AlphaVantage struct {
	Key string
}

func NewAlphaVantage(key string) *AlphaVantage {
	return &AlphaVantage{
		Key: key,
	}
}

// GetStocks returns stocks price variation within a given time frame
// TODO add some logging
// TODO parse body to Stocktiers by Time
func (a *AlphaVantage) GetStocks(stock string, timeFrame string) (map[time.Time]models.StockTier, error) {
	url := buildURL(stock, timeFrame, a.Key)
	body, err := makeHttpCall(url)
	if err != nil {
		return nil, ErrCantGetStockPricesFromAlphaVantage
	}

	fmt.Print(string(body))
	return nil, nil
}

func buildURL(stock string, timeFrame string, key string) string {
	return fmt.Sprintf(URL, stock, timeFrame, key)
}

// makeHttpCall tries to contact AlphaVantage endpoint
// to retrieve stocks price
// TODO add log  with request execution time
// TODO check http status code
func makeHttpCall(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * HTTP_TIMEOUT_SECS,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
