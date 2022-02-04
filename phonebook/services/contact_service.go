package services

import (
	// "time"

	"github.com/seyedmo30/phonebook_api/phonebook/models"
)

type Contact struct {
	ID     uint
	Name   string
	Number string
	Image  string
}

func Add(name *string, number *string) error {

	contact := new(Contact)
	contact.Name = *name
	contact.Number = *number
	obj := models.Contact{Name: contact.Name, Number: contact.Number}
	err := obj.AddContact()
	return err
}
