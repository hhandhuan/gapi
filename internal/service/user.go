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

// HandleRegister 处理用户注册
func (s *userService) HandleRegister(req *entity.UserRegisterRequest) error {
	var user *model.Users
	err := model.User().Where("username = ?", req.Username).Limit(1).Find(&user).Error
	if err != nil {
		return errors.New("user register error")
	}

	if !user.IsEmpty() && !user.IsDelete() {
		return errors.New("username already exists")
	}

	password, err := utils.GenerateFromPassword(req.Password)
	if err != nil {
		return errors.New("user register error")
	}

	err = model.User().Create(&model.Users{Username: req.Username, Password: password}).Error
	if err != nil {
		return errors.New("user register error")
	}

	return nil
}

// HandleLogin 处理用户登录
func (s *userService) HandleLogin(req *entity.UserLoginRequest) (res *gin.H, err error) {
	var user *model.Users
	err = model.User().Where("username", req.Username).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.IsDelete() {
		return nil, errors.New("username or password error")
	}

	if !utils.CompareHashAndPassword(req.Password, user.Password) {
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
	expiredUnix := s.context.jwtClaim.ExpiresAt.Unix()
	currTimeUnix := jwt.NewNumericDate(time.Now()).Unix()

	seconds := expiredUnix - currTimeUnix
	if seconds <= 0 {
		return nil
	}

	// 令牌 ID 对应默认值
	defaultVal := 1
	// 将旧令牌加入到黑名单中
	cmd := redis.DB.Set(context.Background(), s.context.jwtClaim.ID, defaultVal, time.Second*time.Duration(seconds))
	if v, err := cmd.Result(); err != nil || v != "Ok" {
		return errors.New("logout error")
	}

	return nil
}
