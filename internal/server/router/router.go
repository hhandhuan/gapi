package router

import (
	"gapi/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(engine *gin.Engine) {
	engine.NoRoute(handler.Base.NoRoute)
	engine.NoMethod(handler.Base.NoMethod)

	group := engine.Group("api")
	{
		group.POST("login", handler.User.Login)
		group.Use(token())
		{
			group.GET("curr-user", handler.User.CurrUser)
			group.POST("logout", handler.User.Logout)
		}
	}
}
