package user

import (
	"Project-REST-API/entities"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequestFormat struct {
	Nama     string `json:"nama" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}
type GetUsersResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}
type GetUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
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
	Data    entities.User `json:"data"`
}
type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
