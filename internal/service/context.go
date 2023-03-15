package service

import (
	"gapi/internal/consts"
	"gapi/internal/model"
	"gapi/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func (s *contextService) getJwtClaim() *jwt.RegisteredClaims {
	claim, _ := s.context.Get(consts.JwtClaimKey)
	return claim.(*jwt.RegisteredClaims)
}

func (s *contextService) currUser() *model.Users {
	var user *model.Users
	model.User().Where("id", s.getJwtClaim().Subject).Limit(1).Find(&user)
	return user
}
