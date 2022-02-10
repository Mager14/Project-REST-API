package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string

	Task []Task `gorm:"ForeignKey:User_ID"`
}
