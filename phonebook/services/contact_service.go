package services

import (
	"time"

	"github.com/seyedmo30/phonebook_api/phonebook/models"
)


type Contact struct {
	ID        uint           
	Name   string 
	Number string 
	Image  string 
	
}

func (con *Contact) Add() error {
	return models.AddContact(con.Name, con.Number, con.Image)
}
