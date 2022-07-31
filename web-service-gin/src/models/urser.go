package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
	Phone string
	// CreatedAt time.Time      `json:"created_at" form:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	// DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" form:"deleted_at"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Table("user").Find(&User).Error
	fmt.Println(err)
	if err != nil {
		return err

	}
	return nil
}

func PostUsers(db *gorm.DB, User *User) (err error) {
	err = db.Save(User).Error
	fmt.Println(err)
	if err != nil {
		return err

	}
	return nil
}
