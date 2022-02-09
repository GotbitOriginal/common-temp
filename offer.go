package common

import (
	"sort"

	"github.com/shopspring/decimal"
)

type Offer struct {
	Amount decimal.Decimal `json:"amount"`
	Price  decimal.Decimal `json:"price"`
}

var _ sort.Interface = OffersAsc{}

type OffersAsc []Offer

func (offers OffersAsc) Len() int {
	return len(offers)
}

func (offers OffersAsc) Less(i, j int) bool {
	return offers[i].Price.LessThanOrEqual(offers[j].Price)
}

func (offers OffersAsc) Swap(i, j int) {
	offers[i], offers[j] = offers[j], offers[i]
}

var _ sort.Interface = OffersAsc{}

type OffersDesc []Offer

func (offers OffersDesc) Len() int {
	return len(offers)
}

func (offers OffersDesc) Less(i, j int) bool {
	return offers[i].Price.GreaterThan(offers[j].Price)
}

func (offers OffersDesc) Swap(i, j int) {
	offers[i], offers[j] = offers[j], offers[i]
}
