package routes

import (
	"Project-REST-API/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {
	e.GET("users", uc.Get())
	e.GET("users/:id", uc.GetById())
	e.POST("users", uc.Insert())
	e.PUT("users/:id", uc.Update())
	e.DELETE("users/:id", uc.Delete())
}
