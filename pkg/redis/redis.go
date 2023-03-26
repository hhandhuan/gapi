package redis

import (
	"context"
	"fmt"
	"gapi/pkg/conf"
	"github.com/go-redis/redis/v8"
	"log"
)

var DB *redis.Client

func init() {
	DB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.GetConfig().Redis.Host, conf.GetConfig().Redis.Port),
		Password: conf.GetConfig().Redis.Pass,
		DB:       conf.GetConfig().Redis.DB,
	})
	str, err := DB.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Fatalf("redis connect ping failed, err: %v", err)
	}
}
