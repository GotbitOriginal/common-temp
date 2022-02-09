package grpc

import (
	"github.com/gotbitoriginal/commontemp"
	"github.com/shopspring/decimal"
)

func (source *FeeRatios) WriteTo(target *commontemp.FeeRatios) *commontemp.FeeRatios {
	if source == nil || target == nil {
		return target
	}

	target.Maker, _ = decimal.NewFromString(source.Maker)
	target.Taker, _ = decimal.NewFromString(source.Taker)

	return target
}

func (target *FeeRatios) ReadFrom(source *commontemp.FeeRatios) *FeeRatios {
	if target == nil || source == nil {
		return target
	}

	target.Maker = source.Maker.String()
	target.Taker = source.Taker.String()

	return target
}
