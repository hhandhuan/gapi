package router

import (
	"context"
	"gapi/internal/consts"
	"gapi/internal/utils"
	"gapi/pkg/conf"
	"gapi/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
)

func token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := utils.NewResponse(ctx)

		// 请求头解析令牌
		parts := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(parts) <= 1 || parts[0] != "Bearer" {
			response.WithCode(consts.ErrAuthCode).WithMsg(consts.NoRouteErr).JsonOutput()
			ctx.Abort()
			return
		}

		// 解析令牌
		claim, err := utils.NewJwt(conf.GetConfig().Jwt).ParseJwtToken(parts[1])
		if err != nil {
			response.WithCode(consts.ErrAuthCode).WithMsg(consts.NoAuthFailedErr).JsonOutput()
			ctx.Abort()
			return
		}

		val := redis.Client.Get(context.Background(), claim.ID).Val()
		// 检查令牌是否存在黑名单中
		if len(val) > 0 {
			response.WithCode(consts.ErrAuthCode).WithMsg(consts.NoAuthFailedErr).JsonOutput()
			ctx.Abort()
			return
		}

		ctx.Set(consts.JwtClaimKey, claim)
		ctx.Next()
	}
}
