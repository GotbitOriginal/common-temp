package grpc

import (
	"github.com/gotbitoriginal/common"
)

func (source *OrderBook) WriteTo(target *common.OrderBook) *common.OrderBook {
	if source == nil || target == nil {
		return target
	}

	target.Asks = make([]common.Offer, len(source.Asks))
	for i, ask := range source.Asks {
		ask.WriteTo(&target.Asks[i])
	}

	target.Bids = make([]common.Offer, len(source.Bids))
	for i, bid := range source.Bids {
		bid.WriteTo(&target.Bids[i])
	}

	return target
}

func (target *OrderBook) ReadFrom(source *common.OrderBook) *OrderBook {
	if source == nil || target == nil {
		return target
	}

	target.Asks = make([]*Offer, len(source.Asks))
	for i, ask := range source.Asks {
		target.Asks[i] = new(Offer).ReadFrom(&ask)
	}

	target.Bids = make([]*Offer, len(source.Bids))
	for i, bid := range source.Bids {
		target.Bids[i] = new(Offer).ReadFrom(&bid)
	}

	return target

}
