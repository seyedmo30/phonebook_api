package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:128"`
	Email    string `gorm:"uniqueIndex;size:128"`
	Image    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
}

func Get_user_by_username(username *string) (User, error) {
	var user User
	print(&username)
	print(*username)
	err := conn.QueryRow(context.Background(), "select username,email from users where username=$1 ", &username).Scan(&user.Username, &user.Email)

	if err != nil {
		return user, err
	}
	return user, err
}
func Exist_user_by_username(username *string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ?", "hp").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
