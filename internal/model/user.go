package model

import (
	"gapi/pkg/mysql"

	"gorm.io/gorm"
)

type Users struct {
	Base
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
}

func User() *gorm.DB {
	return mysql.GetInstance().Model(&Users{})
}

func (u *Users) GetTableName() string {
	return "users"
}
