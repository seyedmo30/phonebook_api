package models

import (
	// "gorm.io/gorm"
)

type Contact struct {
	ID        uint           `gorm:"primaryKey"`
	Name   string `gorm:"not null;size:128"`
	Number string `gorm:"not null;size:12"`
	Image  string `gorm:"size:256"`
	// Assign User   `gorm:"constraint:OnDelete:SET NULL;"`
}

func AddContact(name string, number string) error {
	contact := Contact{
		Name:      name,
		Number:     number,
	}
	if err := db.Create(&contact).Error; err != nil {
		return err
	}

	return nil
}
