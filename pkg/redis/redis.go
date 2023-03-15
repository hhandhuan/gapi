package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"zhengze/pkg/conf"
)

var Client *redis.Client

func InitClient(conf *conf.Redis) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, gconv.Int(conf.Port)),
		Password: conf.Pass,
		DB:       conf.DB,
		PoolSize: 10,
	})
	str, err := Client.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Printf("redis connect ping failed, err: %v", err)
	}
}
