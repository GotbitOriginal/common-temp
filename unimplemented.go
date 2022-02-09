package commontemp

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

var (
	errUnimplemented = errors.New("unimplemented")
)

var _ Exchange = (*UnimplementedExchange)(nil)

type UnimplementedExchange struct {
}

func (*UnimplementedExchange) GetPairs() ([]Pair, error) {
	return nil, errUnimplemented
}

var _ Markets = (*UnimplementedMarkets)(nil)

type UnimplementedMarkets struct {
}

func (*UnimplementedMarkets) GetPrecision(Pair) (*Precision, error) {
	return nil, errUnimplemented
}

func (*UnimplementedMarkets) GetOrderBook(Pair) (*OrderBook, error) {
	return nil, errUnimplemented
}

var _ Account = (*UnimplementedAccount)(nil)

type UnimplementedAccount struct {
}

func (*UnimplementedAccount) GetBalances() (map[Currency]Balance, error) {
	return nil, errUnimplemented
}

func (*UnimplementedAccount) GetFeeRatios() (*FeeRatios, error) {
	return nil, errUnimplemented
}

func (*UnimplementedAccount) Withdraw(currency Currency, amount decimal.Decimal, address, chain string) (id string, err error) {
	return "", errUnimplemented
}

var _ Trader = (*UnimplementedTrader)(nil)

type UnimplementedTrader struct {
}

func (*UnimplementedTrader) Place(order Order) (OrderID, error) {
	return OrderID(""), errUnimplemented
}

func (*UnimplementedTrader) PlaceMany(orders ...Order) ([]OrderID, error) {
	return nil, errUnimplemented
}

func (*UnimplementedTrader) PlaceMarket(order Order) (OrderID, error) {
	return OrderID(""), errUnimplemented
}

func (*UnimplementedTrader) Cancel(Pair, OrderID) error {
	return errUnimplemented
}

func (*UnimplementedTrader) CancelMany(Pair, ...OrderID) error {
	return errUnimplemented
}

func (*UnimplementedTrader) CancelAll(Pair) error {
	return errUnimplemented
}

func (*UnimplementedTrader) GetOrder(Pair, OrderID) (*Order, error) {
	return nil, errUnimplemented
}

func (*UnimplementedTrader) GetOpenOrders(Pair) ([]Order, error) {
	return nil, errUnimplemented
}

var _ Prices = (*UnimplementedPrices)(nil)

type UnimplementedPrices struct {
}

func (*UnimplementedPrices) GetLastPrice(Pair) (decimal.Decimal, error) {
	return decimal.Zero, errUnimplemented
}

func (*UnimplementedPrices) GetCandles(pair Pair, window Window) ([]Candle, error) {
	return nil, errUnimplemented
}

func (*UnimplementedPrices) GetLastTrades(pair Pair) ([]Trade, error) {
	return nil, errUnimplemented
}

var _ APIClientBase = (*UnimplementedAPIClientBase)(nil)

type UnimplementedAPIClientBase struct {
}

func (*UnimplementedAPIClientBase) Init(...Option) error {
	return errUnimplemented
}

func (*UnimplementedAPIClientBase) Refresh() error {
	return errUnimplemented
}

func (*UnimplementedAPIClientBase) GetRefreshPolicy() *RefreshPolicy {
	return &RefreshPolicy{
		ShouldRefresh: false,
		RefreshDelay:  time.Microsecond,
		RetryCount:    12,
	}
}

func (*UnimplementedAPIClientBase) GetClientID() string {
	return ""
}

type UnimplementedAPIClient struct {
	UnimplementedExchange
	UnimplementedMarkets
	UnimplementedAccount
	UnimplementedTrader
	UnimplementedPrices
	UnimplementedAPIClientBase
}
