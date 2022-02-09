package task

import (
	"Project-REST-API/entities"

	"gorm.io/gorm"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type RegisterTaskRequestFormat struct {
	Nama       string `json:"nama" form:"nama"`
	Priority   int    `json:"priority" form:"priority"`
	User_ID    int    `json:"user_id" form:"user_id"`
	Project_ID int    `json:"project_id" form:"project_id"`
}
type CompletedTaskRequestFormat struct {
	Nama       string `json:"nama" form:"nama"`
	Priority   int    `json:"priority" form:"priority"`
	User_ID    int    `json:"user_id" form:"user_id"`
	Project_ID int    `json:"project_id" form:"project_id"`
}
type ReopenTaskRequestFormat struct {
	Nama       string `json:"nama" form:"nama"`
	Priority   int    `json:"priority" form:"priority"`
	User_ID    int    `json:"user_id" form:"user_id"`
	Project_ID int    `json:"project_id" form:"project_id"`
}
type UpdateTaskRequestFormat struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Priority int    `json:"priority" form:"priority"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
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
