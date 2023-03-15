package handler

import (
	"gapi/internal/consts"
	"gapi/internal/entity"
	"gapi/internal/service"
	"gapi/internal/utils"

	"github.com/gin-gonic/gin"
)

var User = &user{}

type user struct{}

// Login 处理用户登录
func (h *user) Login(ctx *gin.Context) {
	response := utils.NewResponse(ctx)

	var request entity.UserLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.WithCode(consts.ErrParamCode).WithMsg(err).JsonOutput()
		return
	}

	data, err := service.NewUserService(ctx).HandleLogin(&request)
	if err != nil {
		response.WithCode(consts.ErrInternalCode).WithMsg(err).JsonOutput()
		return
	}

	response.WithData(data).JsonOutput()
}

// CurrUser 获取登录信息
func (h *user) CurrUser(ctx *gin.Context) {
	response := utils.NewResponse(ctx)

	currUser := service.NewUserService(ctx).GetCurrUser()

	response.WithData(currUser).JsonOutput()
}
