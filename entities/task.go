package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Nama       string
	Priority   int
	User_ID    int
	Project_ID int
}
