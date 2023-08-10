package repository

import (
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
