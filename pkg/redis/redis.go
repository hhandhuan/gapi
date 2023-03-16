package redis

import (
	"context"
	"fmt"
	"gapi/pkg/conf"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
)

var Client *redis.Client

func InitClient(conf *conf.Redis) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, gconv.Int(conf.Port)),
		Password: conf.Pass,
		DB:       conf.DB,
	})
	str, err := Client.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Fatalf("redis connect ping failed, err: %v", err)
	}
}
