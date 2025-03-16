package redis

import (
	"github.com/go-redis/redis/v8"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
)

type cache struct {
	Client *redis.Client
}

func NewRedis() _interface.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "35.219.50.46:6379",
		Password: "root",
		DB:       1,
	})

	return &cache{
		Client: client,
	}
}
