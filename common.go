package commontemp

import (
	"github.com/shopspring/decimal"
)

type Exchange interface {
	GetPairs() ([]Pair, error)
}

type Markets interface {
	GetPrecision(Pair) (*Precision, error)
	GetOrderBook(Pair) (*OrderBook, error)
}

type Account interface {
	GetBalances() (map[Currency]Balance, error)
	GetFeeRatios() (*FeeRatios, error)
	Withdraw(currency Currency, amount decimal.Decimal, address, chain string) (id string, err error)
}

type Trader interface {
	// Place places the order as limit order.
	Place(order Order) (OrderID, error)
	// PlaceMany places the orders as limit orders.
	PlaceMany(orders ...Order) ([]OrderID, error)
	// PlaceMarket places the order as market order.
	PlaceMarket(order Order) (OrderID, error)

	// Cancel cancels the order by its id.
	Cancel(Pair, OrderID) error
	// CancelMany cancels orders by their ids.
	CancelMany(Pair, ...OrderID) error
	// CancelAll cancels all the orders within the specified pair.
	CancelAll(Pair) error

	// GetOrder returns information about the order.
	GetOrder(Pair, OrderID) (*Order, error)
	// GetOpenOrders returns information about open orders within the specified pair.
	GetOpenOrders(Pair) ([]Order, error)
}

type Prices interface {
	GetCandles(pair Pair, window Window) ([]Candle, error)
	GetLastPrice(pair Pair) (decimal.Decimal, error)
	GetLastTrades(pair Pair) ([]Trade, error)
}

type APIClientBase interface {
	Init(...Option) error
	Refresh() error
	GetRefreshPolicy() *RefreshPolicy
	GetClientID() string
}

type APIClient interface {
	Exchange
	Markets
	Account
	Trader
	Prices

	APIClientBase
}

type APIClientFactory func() APIClientBase
