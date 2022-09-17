package services

import (
	// "time"

	"github.com/seyedmo30/phonebook_api/phonebook/entity"
	"github.com/seyedmo30/phonebook_api/phonebook/models"
)

type Contact struct {
	ID     uint
	Name   string
	Number string
	Image  string
}

func AddContact(contact *entity.Contact) error {
	err := models.AddContact(contact)
	return err
}

func SearchContact(name *string) (*[]entity.Contact, error) {

	contacts, err := models.SearchContact(name)

	return contacts, err
}

func GetContact(id *int) (*entity.Contact, error) {

	contact, err := models.GetContact(id)

	return contact, err
}

func ListContact(page *int) (*[]entity.Contact, error) {

	contacts, err := models.ListContact(page)

	return contacts, err
}
