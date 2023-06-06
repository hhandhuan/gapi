package main

import (
	"context"
	"gapi/internal/server"
	"gapi/pkg/conf"
	"gapi/pkg/logger"
	"gapi/pkg/mysql"
	"gapi/pkg/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config := conf.Initialize()
	logger.Initialize(config.Logger)
	mysql.Initialize(config.Mysql)
	redis.Initialize(config.Redis)
}

func main() {
	httpServer := server.NewServer()

	httpServer.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.GetConfig().System.ShutdownWaitTime))
	defer cancel()

	err := httpServer.Stop(ctx)
	if err != nil {
		logger.GetInstance().Error().Err(err)
	}

	select {
	case <-ctx.Done():
		logger.GetInstance().Info().Msg("api service is down")
		return
	}
}
