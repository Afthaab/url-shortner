package domain

import "time"

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL           string        `json:"url"`
	CustomShort   string        `json:"short"`
	Expiry        time.Duration `json:"expiry"`
	RateRemaining int           `json:"rate_limit"`
	RateLimiting  time.Duration `json:"rate_limit_reset"`
}
