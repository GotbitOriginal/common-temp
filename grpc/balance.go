package grpc

import (
	"github.com/gotbitoriginal/commontemp"
	"github.com/shopspring/decimal"
)

func (source *Balance) WriteTo(target *commontemp.Balance) *commontemp.Balance {
	if source == nil || target == nil {
		return target
	}

	target.Free, _ = decimal.NewFromString(source.Free)
	target.Locked, _ = decimal.NewFromString(source.Locked)

	return target
}

func (target *Balance) ReadFrom(source *commontemp.Balance) *Balance {
	if target == nil || source == nil {
		return target
	}

	target.Free = source.Free.String()
	target.Locked = source.Locked.String()

	return target
}
