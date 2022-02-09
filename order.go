package common

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID               OrderID         `json:"id"`
	Side             Side            `json:"side"`       // Side specifies the side of the order (see SideBuy, SideSell).
	Pair             Pair            `json:"pair"`       // Pair specifies currency pair of the order.
	Amount           decimal.Decimal `json:"amount"`     // Amount specifies full amount of base currency.
	AmountFilled     decimal.Decimal `json:"filled"`     // AmountFilled specifies filled amount of base currency.
	Price            decimal.Decimal `json:"price"`      // Price contans max-buy/min-sell amount of quoted currency to give/take for each base currency unit.
	PriceFillAgerage decimal.Decimal `json:"average"`    // PriceFillAgerage contains average price of the already filled volume.
	Fee              decimal.Decimal `json:"fee"`        // Fee is amount of quote currency spent to pay fees.
	Status           Status          `json:"status"`     // Status specifies the status of the order (see StatusNew, StatusOpen, StatusClosed).
	Time             time.Time       `json:"time"`       // Time specifies when the order was created.
	FillOrKill       bool            `json:"fillOrKill"` // FillOrKill specifies whether the order is 'fill or kill' order.
}

func (o *Order) GetBaseAmount() decimal.Decimal {
	return o.Amount
}

func (o *Order) GetQuoteAmount() decimal.Decimal {
	return decimal.Sum(
		o.AmountFilled.Mul(o.PriceFillAgerage),
		o.Amount.Sub(o.AmountFilled).Mul(o.Price),
	)
}

func (o *Order) GetBaseAmountFilled() decimal.Decimal {
	return o.AmountFilled
}

func (o *Order) GetQuoteAmountFilled() decimal.Decimal {
	return o.AmountFilled.Mul(o.PriceFillAgerage)
}

func (o *Order) GetFee() decimal.Decimal {
	return o.Fee
}
