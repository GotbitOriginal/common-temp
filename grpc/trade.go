package grpc

import (
	"time"

	"github.com/gotbitoriginal/commontemp"
	"github.com/shopspring/decimal"
)

func (source *Trade) WriteTo(target *commontemp.Trade) *commontemp.Trade {
	if source == nil || target == nil {
		return target
	}

	target.Time = time.Unix(source.Time, 0)
	target.Side = commontemp.Side(source.Side)
	target.Amount, _ = decimal.NewFromString(source.Amount)
	target.Price, _ = decimal.NewFromString(source.Price)

	return target
}

func (target *Trade) ReadFrom(source *commontemp.Trade) *Trade {
	if source == nil || target == nil {
		return target
	}

	target.Time = source.Time.Unix()
	target.Side = string(source.Side)
	target.Amount = source.Amount.String()
	target.Price = source.Price.String()

	return target
}
