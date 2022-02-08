package routes

import (
	"Project-REST-API/delivery/controllers/task"
	"Project-REST-API/delivery/controllers/user"
	"Project-REST-API/middlewares"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, tc *task.TaskController) {
	e.POST("users/register", uc.UserRegister())
	e.POST("users/login", uc.Login())
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("users", uc.Get())
	eAuth.GET("users/:id", uc.GetById())
	eAuth.PUT("users/:id", uc.Update())
	eAuth.DELETE("users/:id", uc.Delete())

	e.POST("task/register", tc.TaskRegister())
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("task", tc.Get())
	eAuth.GET("task/:id", tc.GetById())
	eAuth.PUT("task/:id", tc.Update())
	eAuth.DELETE("task/:id", tc.Delete())
}
