package commontemp

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
)

func marshal(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func log(t *testing.T, v ...interface{}) {
	t.Log(v...)
}

func TestJson(t *testing.T) {
	order := Order{
		ID:               "12",
		Side:             SideBuy,
		Pair:             Pair{Base: "BTC", Quote: "USDT"},
		Amount:           decimal.NewFromFloat(1232),
		AmountFilled:     decimal.NewFromFloat(1),
		PriceFillAgerage: decimal.NewFromFloat(122),
		Price:            decimal.NewFromFloat(123),
	}
	log(t, marshal(&order))
	log(t, marshal(order.GetBaseAmount()))
	log(t, marshal(order.GetBaseAmountFilled()))

	log(t, marshal(order.GetQuoteAmount()))
	log(t, marshal(order.GetQuoteAmountFilled()))

	log(t, marshal(order.Price))
	log(t, marshal(order.PriceFillAgerage))

}
