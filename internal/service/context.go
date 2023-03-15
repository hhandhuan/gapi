package service

import (
	"zhengze/internal/consts"
	"zhengze/internal/model"
	"zhengze/pkg/conf"

	"github.com/gin-gonic/gin"
)

type contextService struct {
	context *gin.Context
	config  *conf.Config
}

func NewContextService(ctx *gin.Context) *contextService {
	return &contextService{
		context: ctx,
		config:  conf.GetConfig(),
	}
}

func (s *contextService) getJwtSub() string {
	uid, _ := s.context.Get(consts.JwtSubKey)
	return uid.(string)
}

func (s *contextService) currUser() *model.Users {
	var user *model.Users
	model.User().Where("id", s.getJwtSub()).Limit(1).Find(&user)
	return user
}
