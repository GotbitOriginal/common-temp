package commontemp

type OrderBook struct {
	Asks []Offer `json:"asks"`
	Bids []Offer `json:"bids"`
}
