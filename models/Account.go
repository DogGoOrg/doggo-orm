package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email       string `gorm:"type:varchar(200);UNIQUE"`
	Password    string `gorm:"type:varchar(200);"`
	FullName    string `gorm:"type:varchar(200);"`
	IsActivated bool   `gorm:"type:bool;default:false"`
}
