package entity

type UserLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
