package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string `gorm:"uniqueIndex;size:128"`
	email    string `gorm:"uniqueIndex;size:128"`
	image    string `gorm:"size:256"`
	password string `gorm:"size:256"`
}
