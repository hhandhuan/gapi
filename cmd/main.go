package main

import (
	"context"
	"gapi/internal/server"
	"gapi/pkg/conf"
	"gapi/pkg/consts"
	"gapi/pkg/mysql"
	"gapi/pkg/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	config := conf.GetConfig()
	mysql.InitClient(config.Mysql)
	redis.InitClient(config.Redis)
}

func main() {
	svr := server.NewServer()

	go func() {
		if err := svr.Run(); err != nil {
			log.Printf("server start error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), consts.ServerShutdownWaitTime)
	defer cancel()

	if err := svr.Stop(ctx); err != nil {
		log.Println("stop server error: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("server exiting")
}
