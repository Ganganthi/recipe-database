package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
}

type Ingredient struct {
	gorm.Model
	Name string
}
