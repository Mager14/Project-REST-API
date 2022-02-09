package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Nama string
}
