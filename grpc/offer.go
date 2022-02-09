package grpc

import (
	"github.com/gotbitoriginal/commontemp"
	"github.com/shopspring/decimal"
)

func (source *Offer) WriteTo(target *commontemp.Offer) *commontemp.Offer {
	if source == nil || target == nil {
		return target
	}

	target.Amount, _ = decimal.NewFromString(source.Amount)
	target.Price, _ = decimal.NewFromString(source.Price)

	return target
}

func (target *Offer) ReadFrom(source *commontemp.Offer) *Offer {
	if target == nil || source == nil {
		return target
	}

	target.Amount = source.Amount.String()
	target.Price = source.Price.String()

	return target
}
