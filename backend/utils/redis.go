package utils

import{
	"context"
	"github.com/go-redis/redis/v8"
}

var Ctx = context.Background()
var RedisClient *redis.Client

func InitRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}