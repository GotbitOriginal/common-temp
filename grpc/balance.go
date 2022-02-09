package grpc

import (
	"github.com/gotbitoriginal/common"
	"github.com/shopspring/decimal"
)

func (source *Balance) WriteTo(target *common.Balance) *common.Balance {
	if source == nil || target == nil {
		return target
	}

	target.Free, _ = decimal.NewFromString(source.Free)
	target.Locked, _ = decimal.NewFromString(source.Locked)

	return target
}

func (target *Balance) ReadFrom(source *common.Balance) *Balance {
	if target == nil || source == nil {
		return target
	}

	target.Free = source.Free.String()
	target.Locked = source.Locked.String()

	return target
}
