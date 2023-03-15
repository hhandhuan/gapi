package router

import (
	"strings"
	"zhengze/internal/consts"
	"zhengze/internal/utils"
	"zhengze/pkg/conf"

	"github.com/gin-gonic/gin"
)

func token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := utils.NewResponse(ctx)

		parts := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(parts) <= 1 || parts[0] != "Bearer" {
			response.WithCode(consts.ErrAuthCode).WithMsg(consts.NoRouteErr).JsonOutput()
			ctx.Abort()
			return
		}

		claim, err := utils.NewJwt(conf.GetConfig().Jwt).ParseJwtToken(parts[1])
		if err != nil {
			response.WithCode(consts.ErrAuthCode).WithMsg(consts.NoAuthFailedErr).JsonOutput()
			ctx.Abort()
			return
		}

		ctx.Set(consts.JwtSubKey, claim.Subject)
		ctx.Next()
	}
}
