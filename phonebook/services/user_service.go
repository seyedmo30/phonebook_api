package services

import (
	"time"

	"github.com/seyedmo30/phonebook_api/phonebook/models"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Image     string
	Password  string
}

func (u *User) Exist_by_username() (bool, error) {
	return models.Exist_user_by_username(&u.Username)
}
