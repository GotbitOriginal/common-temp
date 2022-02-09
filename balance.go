package common

import "github.com/shopspring/decimal"

type Balance struct {
	Free   decimal.Decimal `json:"free"`
	Locked decimal.Decimal `json:"locked"`
}
