package grpc

import (
	"github.com/gotbitoriginal/commontemp"
)

func (source *OrderBook) WriteTo(target *commontemp.OrderBook) *commontemp.OrderBook {
	if source == nil || target == nil {
		return target
	}

	target.Asks = make([]commontemp.Offer, len(source.Asks))
	for i, ask := range source.Asks {
		ask.WriteTo(&target.Asks[i])
	}

	target.Bids = make([]commontemp.Offer, len(source.Bids))
	for i, bid := range source.Bids {
		bid.WriteTo(&target.Bids[i])
	}

	return target
}

func (target *OrderBook) ReadFrom(source *commontemp.OrderBook) *OrderBook {
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
