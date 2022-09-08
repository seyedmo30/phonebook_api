package entity

type Contact struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null;size:128"`
	Number int16  `json:"number" gorm:"not null;size:12"`
	Image  string `json:"image" gorm:"size:256"`
}
