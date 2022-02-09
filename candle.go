package commontemp

import (
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Time   time.Time       `json:"time"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Volume decimal.Decimal `json:"amount"`
}

var _ sort.Interface = Candles{}

type Candles []Candle

func (candles Candles) Len() int {
	return len(candles)
}

func (candles Candles) Less(i, j int) bool {
	return candles[j].Time.After(candles[i].Time)
}

func (candles Candles) Swap(i, j int) {
	candles[i], candles[j] = candles[j], candles[i]
}
