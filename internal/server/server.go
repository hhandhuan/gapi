package server

import (
	"context"
	"fmt"
	"gapi/internal/server/router"
	"gapi/pkg/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	server *http.Server
}

func NewServer() *server {
	gin.SetMode(conf.GetConfig().System.Env)

	engine := gin.New()

	router.RegisterRouter(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.GetConfig().System.Addr),
		Handler: engine,
	}

	return &server{server: srv}
}

// Init 初始化服务
func (s *server) init() {
}

// Run 运行服务
func (s *server) Run() {
	go func() {
		s.server.ListenAndServe()
	}()
}

// Stop 停止服务
func (s *server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
