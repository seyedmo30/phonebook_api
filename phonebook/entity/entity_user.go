package entity

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;size:128"`
	Email    string `json:"email" gorm:"uniqueIndex;size:128"`
	Image    string `json:"image" gorm:"size:256"`
	Password string `json:"password" gorm:"size:256"`
}
