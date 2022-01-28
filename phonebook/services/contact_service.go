package services

import (
	// "time"

	"fmt"

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

	fmt.Println("ser \n", contact.Name)
	fmt.Println("ser \n", contact.Number)

	err := models.AddContact(&contact.Name, &contact.Number)
	return err
}
