package grpc

import "github.com/gotbitoriginal/commontemp"

func (source *Precision) WriteTo(target *commontemp.Precision) *commontemp.Precision {
	if source == nil || target == nil {
		return target
	}

	target.Amount = source.Amount
	target.Price = source.Price

	return target
}

func (target *Precision) ReadFrom(source *commontemp.Precision) *Precision {
	if source == nil || target == nil {
		return target
	}

	target.Amount = source.Amount
	target.Price = source.Price

	return target
}
