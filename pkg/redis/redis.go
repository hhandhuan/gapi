package redis

import (
	"context"
	"fmt"
	"gapi/pkg/conf"
	"github.com/go-redis/redis/v8"
	"log"
)

var instance *redis.Client

func GetInstance() *redis.Client {
	return instance
}

func Initialize(cf *conf.Redis) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Pass,
		DB:       cf.DB,
	})
	str, err := client.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Fatalf("redis connect ping failed, err: %v", err)
	}
	instance = client
}
