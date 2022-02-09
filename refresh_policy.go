package commontemp

import "time"

type RefreshPolicy struct {
	ShouldRefresh bool          `json:"shouldRefresh"`
	RefreshDelay  time.Duration `json:"delay"`
	RetryCount    int           `json:"retryCount"`
}
