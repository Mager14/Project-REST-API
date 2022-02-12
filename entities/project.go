package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Nama string
	Task []Task `gorm:"ForeignKey:Project_ID"`
}
