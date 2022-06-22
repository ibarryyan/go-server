package config

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

var RedisUrl, RedisPwd string
var RedisDb int

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     RedisUrl,
		Password: RedisPwd,
		DB:       RedisDb,
	})
}
