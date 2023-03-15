package server

import (
	"context"
	"net/http"
	"zhengze/internal/server/router"
	"zhengze/pkg/conf"

	"github.com/gin-gonic/gin"
)

type server struct {
	server *http.Server
}

func NewServer() *server {
	gin.SetMode(conf.GetConfig().System.Env)

	engine := gin.New()
	router.RegisterRouter(engine)

	httpServer := &http.Server{Addr: ":8080", Handler: engine}

	return &server{server: httpServer}
}

// Init 初始化服务
func (s *server) init() {
}

// Run 运行服务
func (s *server) Run() error {
	return s.server.ListenAndServe()
}

// Stop 停止服务
func (s *server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
