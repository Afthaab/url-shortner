package interfaces

import "time"

type UrlRepository interface {
	FindTheIp(ip string) (string, error)
	SaveTheIp(ip string) error
	CheckTheTime(ip string) (time.Duration, error)
	FindTheURL(url string) (string, error)
	SetTheUrl(url string, id string, expiry time.Duration) error
}
