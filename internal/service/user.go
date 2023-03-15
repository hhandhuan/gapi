package service

import (
	"context"
	"errors"
	"gapi/internal/entity"
	"gapi/internal/model"
	"gapi/internal/utils"
	"gapi/pkg/redis"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func NewUserService(ctx *gin.Context) *userService {
	return &userService{context: NewContextService(ctx)}
}

type userService struct {
	context *contextService
}

// HandleLogin 处理用户登录
func (s *userService) HandleLogin(req *entity.UserLoginRequest) (res *gin.H, err error) {
	var user *model.Users
	err = model.User().Where("username", req.Username).Scopes(model.UnTrash).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID <= 0 {
		return nil, errors.New("username or password error")
	}

	if utils.MD5(req.Password) != user.Password {
		return nil, errors.New("username or password error")
	}

	token, err := utils.NewJwt(s.context.config.Jwt).CreateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, errors.New("user login error")
	}

	data := &gin.H{"token": token, "ttl": s.context.config.Jwt.Ttl, "type": "Bearer"}

	return data, nil
}

// GetCurrUser 获取当前登录的用户
func (s *userService) GetCurrUser() *model.Users {
	return s.context.currUser()
}

// HandleLogout 处理用户登录
func (s *userService) HandleLogout() error {
	seconds := (*s.context.getJwtClaim().ExpiresAt).Unix() - jwt.NewNumericDate(time.Now()).Unix()

	cmd := redis.Client.Set(context.Background(), s.context.getJwtClaim().ID, 1, time.Second*time.Duration(seconds))
	if v, err := cmd.Result(); err != nil || v != "Ok" {
		return errors.New("server internal error")
	}

	return nil
}
