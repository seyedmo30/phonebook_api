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

func SearchContact(name *string) (*[]entity.Contact, error) {
	var contacts []entity.Contact

	rows, err := conn.Query(context.Background(), "select name,number from contacts where name like $1 ", "%"+*name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number int64
		var name string
		err := rows.Scan(&name, &number)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, entity.Contact{
			Name: name, Number: number,
		})

	}

	return &contacts, nil
}

func GetContact(id *int) (*entity.Contact, error) {
	var contact entity.Contact
	err := conn.QueryRow(context.Background(), "select name,number from contacts where id = $1 ", *id).Scan(&contact.Name, &contact.Number)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
func AddContact(contact *entity.Contact) error {

	_, err := conn.Exec(context.Background(), "insert into contacts (name,number,image) values ($1,$2,$3); ", *&contact.Name, *&contact.Number, *&contact.Image)
	if err != nil {
		return err
	}

	return nil
}

func ListContact(page *int) (*[]entity.Contact, error) {
	var contacts []entity.Contact

	rows, err := conn.Query(context.Background(), "select id,name,number, COALESCE(image, '') as image from contacts limit $1 ", &page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var name string
		var number int64
		var image string
		err := rows.Scan(&id, &name, &number, &image)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, entity.Contact{
			ID: id, Name: name, Number: number, Image: image,
		})

	}

	return &contacts, nil
}
