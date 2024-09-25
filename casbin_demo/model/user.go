package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	UserName  string `gorm:"unique;column:用户名,"`
	NickName  string
	PassWord  string
	Telephone string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
