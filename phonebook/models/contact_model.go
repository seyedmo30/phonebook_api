package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	name   string `gorm:"not null;size:128"`
	number string `gorm:"not null;size:12"`
	image  string `gorm:"size:256"`
	assign User   `gorm:"constraint:OnDelete:SET NULL;"`
}
