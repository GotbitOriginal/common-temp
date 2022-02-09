package grpc

import "github.com/gotbitoriginal/common"

func (source *Pair) WriteTo(target *common.Pair) *common.Pair {
	if source == nil || target == nil {
		return target
	}

	target.Base = common.Currency(source.Base)
	target.Quote = common.Currency(source.Quote)

	return target
}

func (target *Pair) ReadFrom(source *common.Pair) *Pair {
	if source == nil || target == nil {
		return target
	}

	target.Base = string(source.Base)
	target.Quote = string(source.Quote)

	return target
}
