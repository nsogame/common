package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID int

	Username string
	Password string
}
