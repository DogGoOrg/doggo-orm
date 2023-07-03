package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(200);"`
	Species  string `gorm:"type:varchar(200);"`
	Breed    string `gorm:"type:varchar(200);"`
}
