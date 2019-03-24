package client

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/pedrolopesme/sentinel/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	URL               = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%v&interval=%v&outputsize=full&datatype=csv&apikey=%v"
	HTTP_TIMEOUT_SECS = 10
	TIME_LAYOUT       = "2006-01-02 15:04:05"
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
	return parseStocksCSV(body)
}

// GetName returns stock provider name
func (a *AlphaVantage) GetName() string {
	return "AlphaVantage"
}

func buildURL(stock string, timeFrame string, key string) string {
	return fmt.Sprintf(URL, stock, timeFrame, key)
}

// makeHttpCall tries to contact AlphaVantage endpoint
// to retrieve stocks price
// TODO add log  with request execution time
// TODO check http status code
// TODO extract to a http lib?
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

// parseStocksCSV knows how to transform an AlphaVantage response to
// a list of stocks
// TODO add tests
// TODO add logs
// TODO add util funcs for time, number manipulation
func parseStocksCSV(body []byte) (map[time.Time]models.StockTier, error) {
	const (
		TIME_COLUMN = iota
		OPEN_COLUMN
		HIGH_COLUMN
		LOW_COLUMN
		CLOSE_COLUMN
		VOLUME_COLUMN
	)

	var bytesReader = bytes.NewReader(body)
	var csvReader = csv.NewReader(bytesReader)
	data, err := csvReader.ReadAll()
	if err != nil {
		// TODO log real error cause
		return nil, ErrCantParseStockPricesFromAlphaVantage
	}

	var stockTiers = make(map[time.Time]models.StockTier)
	for line, entry := range data {
		if line == 0 {
			continue // skipping header
		}

		stockTime, err := time.Parse(TIME_LAYOUT, entry[TIME_COLUMN])
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		stockVolume, err := strconv.ParseInt(entry[VOLUME_COLUMN], 10, 64)
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		stockOpen, err := strconv.ParseFloat(entry[OPEN_COLUMN], 64)
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		stockHigh, err := strconv.ParseFloat(entry[HIGH_COLUMN], 64)
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		stockLow, err := strconv.ParseFloat(entry[LOW_COLUMN], 64)
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		stockClose, err := strconv.ParseFloat(entry[CLOSE_COLUMN], 64)
		if err != nil {
			// TODO log real error cause
			fmt.Println(err.Error())
			return nil, ErrCantParseStockPricesFromAlphaVantage
		}

		tier := models.StockTier{
			Volume: stockVolume,
			Price: &models.StockPrice{
				Open:  models.Money(stockOpen),
				Close: models.Money(stockClose),
				Low:   models.Money(stockLow),
				High:  models.Money(stockHigh),
			},
		}
		stockTiers[stockTime] = tier
	}

	return stockTiers, nil
}
