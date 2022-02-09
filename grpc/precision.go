package grpc

import "github.com/gotbitoriginal/common"

func (source *Precision) WriteTo(target *common.Precision) *common.Precision {
	if source == nil || target == nil {
		return target
	}

	target.Amount = source.Amount
	target.Price = source.Price

	return target
}

func (target *Precision) ReadFrom(source *common.Precision) *Precision {
	if source == nil || target == nil {
		return target
	}

	target.Amount = source.Amount
	target.Price = source.Price

	return target
}
