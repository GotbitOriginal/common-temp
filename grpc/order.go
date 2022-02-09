package grpc

import (
	"time"

	"github.com/gotbitoriginal/commontemp"
	"github.com/shopspring/decimal"
)

func (source *Order) WriteTo(target *commontemp.Order) *commontemp.Order {
	if source == nil || target == nil {
		return target
	}

	target.ID = commontemp.OrderID(source.Id)
	target.Side = commontemp.Side(source.Side)
	source.Pair.WriteTo(&target.Pair)
	target.Amount, _ = decimal.NewFromString(source.Amount)
	target.AmountFilled, _ = decimal.NewFromString(source.AmountFilled)
	target.Price, _ = decimal.NewFromString(source.Price)
	target.PriceFillAgerage, _ = decimal.NewFromString(source.PriceFillAgerage)
	target.Fee, _ = decimal.NewFromString(source.Fee)
	target.Status = commontemp.Status(source.Status)
	target.Time = time.Unix(source.Time, 0)
	target.FillOrKill = source.FillOrKill

	return target
}

func (target *Order) ReadFrom(source *commontemp.Order) *Order {
	if source == nil || target == nil {
		return target
	}

	target.Id = string(source.ID)
	target.Side = string(source.Side)
	target.Pair = new(Pair).ReadFrom(&source.Pair)
	target.Amount = source.Amount.String()
	target.AmountFilled = source.AmountFilled.String()
	target.Price = source.Price.String()
	target.PriceFillAgerage = source.PriceFillAgerage.String()
	target.Fee = source.Fee.String()
	target.Status = string(source.Status)
	target.Time = source.Time.Unix()
	target.FillOrKill = source.FillOrKill

	return target
}
