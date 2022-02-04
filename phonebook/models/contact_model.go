package models

// "gorm.io/gorm"

type Contact struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null;size:128"`
	Number string `gorm:"not null;size:12"`
	Image  string `gorm:"size:256"`
	// Assign User   `gorm:"constraint:OnDelete:SET NULL;"`
}

func (struct_obj *Contact) AddContact() error {

	if err := db.Create(&struct_obj).Error; err != nil {
		return err
	}

	return nil
}

// func GetListContact(pageNum int, pageSize int, maps interface{}) ([]*Contact, error) {

// }

func (con *Contact) AddContact2() error {
	if err := db.Create(&con).Error; err != nil {
		return err
	}

	return nil
}
