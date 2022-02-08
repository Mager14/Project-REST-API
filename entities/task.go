package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Nama       string
	User_Id    int
	Priority   int
	Project_Id int
}
