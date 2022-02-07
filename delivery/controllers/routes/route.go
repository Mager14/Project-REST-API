package routes

import (
	"Project-REST-API/delivery/controllers/user"
	"Project-REST-API/middlewares"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {
	e.POST("users/register", uc.UserRegister())
	e.POST("users/login", uc.Login())
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("users", uc.Get())
	eAuth.GET("users/:id", uc.GetById())
	eAuth.PUT("users/:id", uc.Update())
	eAuth.DELETE("users/:id", uc.Delete())
}
