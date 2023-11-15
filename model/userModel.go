package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
	Uid      string `json:"uid"`
}
