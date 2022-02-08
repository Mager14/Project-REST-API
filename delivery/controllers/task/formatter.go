package task

import (
	"Project-REST-API/entities"

	"gorm.io/gorm"
)

type RegisterRequestFormat struct {
	Nama     string `json:"nama" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterTaskResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type GetTasksResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.Task `json:"data"`
}

type GetTaskResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type UpdateRequestFormat struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
