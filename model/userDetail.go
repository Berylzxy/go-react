package model

import "time"

type UserDetail struct {
	UserName string    `gorm:"username"`
	Birth    time.Time `gorm:"birth"'`
}
