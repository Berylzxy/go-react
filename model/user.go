package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"id"`
	Userid   int    `gorm:"userid"`
	UserName string `gorm:"username"`
	PassWord string `gorm:"password"`
}
