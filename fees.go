package commontemp

import "github.com/shopspring/decimal"

type FeeRatios struct {
	Maker decimal.Decimal `json:"maker"`
	Taker decimal.Decimal `json:"taker"`
}
