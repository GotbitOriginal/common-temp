package grpc

import (
	"time"

	"github.com/gotbitoriginal/common"
	"github.com/shopspring/decimal"
)

func (source *Candle) WriteTo(target *common.Candle) *common.Candle {
	if source == nil || target == nil {
		return target
	}

	target.Time = time.Unix(source.Time, 0)
	target.Open, _ = decimal.NewFromString(source.Open)
	target.Close, _ = decimal.NewFromString(source.Close)
	target.High, _ = decimal.NewFromString(source.High)
	target.Low, _ = decimal.NewFromString(source.Low)
	target.Volume, _ = decimal.NewFromString(source.Volume)

	return target
}

func (target *Candle) ReadFrom(source *common.Candle) *Candle {
	if source == nil || target == nil {
		return target
	}

	target.Time = source.Time.Unix()
	target.Open = source.Open.String()
	target.Close = source.Close.String()
	target.High = source.High.String()
	target.Low = source.Low.String()
	target.Volume = source.Volume.String()

	return target
}
