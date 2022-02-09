package grpc

import "github.com/gotbitoriginal/commontemp"

func (source *Pair) WriteTo(target *commontemp.Pair) *commontemp.Pair {
	if source == nil || target == nil {
		return target
	}

	target.Base = commontemp.Currency(source.Base)
	target.Quote = commontemp.Currency(source.Quote)

	return target
}

func (target *Pair) ReadFrom(source *commontemp.Pair) *Pair {
	if source == nil || target == nil {
		return target
	}

	target.Base = string(source.Base)
	target.Quote = string(source.Quote)

	return target
}
