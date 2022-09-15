package models

import (
	"context"

	"github.com/seyedmo30/phonebook_api/phonebook/entity"
)

// "gorm.io/gorm"

type Contact struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null;size:128"`
	Number string `gorm:"not null;size:12"`
	Image  string `gorm:"size:256"`
	// Assign User   `gorm:"constraint:OnDelete:SET NULL;"`
}

func SearchContact(name *string) (*entity.Contact, error) {
	var contact entity.Contact
	err := conn.QueryRow(context.Background(), "select name,number from contact where name like $1 ", "%"+*name+"%").Scan(&contact.Name, &contact.Number)

	if err != nil {
		return &contact, err
	}
	return &contact, nil
}

func (struct_obj *Contact) AddContact() error {

	if err := db.Create(&struct_obj).Error; err != nil {
		return err
	}

	return nil
}

// func GetListContact(pageNum int, pageSize int, maps interface{}) ([]*Contact, error) {

// }

func (con *Contact) AddContact2() error {
	if err := db.Create(&con).Error; err != nil {
		return err
	}

	return nil
}
