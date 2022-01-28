package models

import "fmt"

// "gorm.io/gorm"

type Contact struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null;size:128"`
	Number string `gorm:"not null;size:12"`
	Image  string `gorm:"size:256"`
	// Assign User   `gorm:"constraint:OnDelete:SET NULL;"`
}

func AddContact(name *string, number *string) error {
	contact := Contact{
		Name:   *name,
		Number: *number,
	}
	fmt.Println("mod \n", contact.Name)
	fmt.Println("mod \n", contact.Number)

	if err := db.Create(contact).Error; err != nil {
		return err
	}

	return nil
}

func (con *Contact) AddContact2() error {
	if err := db.Create(&con).Error; err != nil {
		return err
	}

	return nil
}
