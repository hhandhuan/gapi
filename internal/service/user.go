package service

import (
	"errors"
	"gapi/internal/entity"
	"gapi/internal/model"
	"gapi/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	usernameOrPasswordError = errors.New("username or password error")
	userLoginError          = errors.New("login error")
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
		return nil, usernameOrPasswordError
	}

	if utils.MD5(req.Password) != user.Password {
		return nil, usernameOrPasswordError
	}

	jwt := utils.NewJwt(s.context.config.Jwt)

	token, err := jwt.CreateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, userLoginError
	}

	data := &gin.H{"token": token, "ttl": s.context.config.Jwt.Ttl, "type": "Bearer"}

	return data, nil
}

// GetCurrUser 获取当前登录的用户
func (s *userService) GetCurrUser() *model.Users {
	return s.context.currUser()
}
