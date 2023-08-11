package repository

import (
	"os"
	"time"

	interfaces "github.com/afthaab/urlshortner/pkg/repository/interface"
	"github.com/go-redis/redis"
)

type urlDataBase struct {
	RDB1 *redis.Client
	RDB2 *redis.Client
}

func NewUrlRepository(rdb1 *redis.Client, RDB2 *redis.Client) interfaces.UrlRepository {
	return &urlDataBase{
		RDB1: rdb1,
		RDB2: RDB2,
	}
}

func (r *urlDataBase) FindTheIp(ip string) (string, error) {
	value, err := r.RDB2.Get(ip).Result()
	return value, err
}

func (r *urlDataBase) SaveTheIp(ip string) error {
	err := r.RDB2.Set(ip, os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	return err
}

func (r *urlDataBase) CheckTheTime(ip string) (time.Duration, error) {
	value, err := r.RDB2.TTL(ip).Result()
	return value, err
}

func (r *urlDataBase) FindTheURL(url string) (string, error) {
	value, err := r.RDB1.Get(url).Result()
	return value, err
}

func (r *urlDataBase) SetTheUrl(url string, id string, expiry time.Duration) error {
	err := r.RDB1.Set(id, url, expiry*3600*time.Second).Err()
	return err
}

func (r *urlDataBase) IncrementTheCounter(key string) *redis.IntCmd {
	return r.RDB2.Incr(key)
}
