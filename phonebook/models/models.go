package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Contact{})
	db.Create(&User{Username: "hp"})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func CloseDB() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
