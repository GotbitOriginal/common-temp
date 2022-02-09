package common

type Currency string

type OrderID string

type Side string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"
)

type Status string

const (
	StatusNew    Status = "new"
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

type Window string

const (
	Min1   Window = "min1"
	Min5   Window = "min5"
	Min15  Window = "min15"
	Min30  Window = "min30"
	Hour1  Window = "hour1"
	Hour12 Window = "hour12"
	Day1   Window = "day1"
	Week1  Window = "week1"
	Month1 Window = "month1"
)

func (w Window) Minutes() int64 {
	switch w {
	case Min1:
		return 1
	case Min5:
		return 5
	case Min15:
		return 15
	case Min30:
		return 30
	case Hour1:
		return 1 * 60
	case Hour12:
		return 12 * 60
	case Day1:
		return 24 * 60
	case Week1:
		return 7 * 24 * 60
	case Month1:
		return 30 * 24 * 60
	default:
		return 5
	}
}
