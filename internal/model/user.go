package model

import (
	"gapi/pkg/mysql"

	"gorm.io/gorm"
)

type Users struct {
	BaseModel
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
}

func User() *gorm.DB {
	return mysql.DB.Model(&Users{})
}
