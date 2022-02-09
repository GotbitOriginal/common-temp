package commontemp

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	Time   time.Time       `json:"time"`
	Side   Side            `json:"side"`
	Amount decimal.Decimal `json:"amount"`
	Price  decimal.Decimal `json:"price"`
}
