package service

import (
	"gapi/internal/consts"
	"gapi/internal/model"
	"gapi/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type contextService struct {
	context  *gin.Context
	config   *conf.Config
	jwtClaim *jwt.RegisteredClaims
}

func NewContextService(ctx *gin.Context) *contextService {
	context := &contextService{
		context: ctx,
		config:  conf.GetConfig(),
	}
	if claim, _ := ctx.Get(consts.JwtClaimKey); claim != nil {
		context.jwtClaim = claim.(*jwt.RegisteredClaims)
	}

	return context
}

func (s *contextService) currUser() *model.Users {
	var user *model.Users
	model.User().Where("id", s.jwtClaim.Subject).Limit(1).Find(&user)
	return user
}
