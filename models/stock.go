package models

// Money extracts currency to its own type
type Money float64

// StockPrice stores the price variation within a range of time
type StockPrice struct {
	Open  Money
	Close Money
	High  Money
	Low   Money
}

// StockTier joins price variation with other important information
type StockTier struct {
	Volume int64
	Price  *StockPrice
}
