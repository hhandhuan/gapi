package handler

import (
	"github.com/gin-gonic/gin"
	"zhengze/internal/consts"
	"zhengze/internal/utils"
)

var Base = &base{}

type base struct{}

func (h *base) NoRoute(ctx *gin.Context) {
	response := utils.NewResponse(ctx)
	response.WithCode(consts.ErrNoRouteCode).WithMsg(consts.NoRouteErr).JsonOutput()
}

func (h *base) NoMethod(ctx *gin.Context) {
	response := utils.NewResponse(ctx)
	response.WithCode(consts.ErrNoMethodCode).WithMsg(consts.NoMethodErr).JsonOutput()
}
