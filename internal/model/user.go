package model

import (
	"zhengze/pkg/mysql"

	"gorm.io/gorm"
)

type Users struct {
	BaseModel
	Username    string `gorm:"column:username" json:"username"`
	Password    string `gorm:"column:password" json:"-"`
	AccessToken string `gorm:"column:access_token" json:"-"`
}

func User() *gorm.DB {
	return mysql.Cli.DB.Model(&Users{})
}
