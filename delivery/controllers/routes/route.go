package routes

import (
	"Project-REST-API/configs"
	"Project-REST-API/delivery/controllers/project"
	"Project-REST-API/delivery/controllers/task"
	"Project-REST-API/delivery/controllers/user"
	"Project-REST-API/middlewares"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, tc *task.TaskController, pc *project.ProjectController) {

	e.POST("users/register", uc.UserRegister())
	e.POST("users/login", uc.Login())
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("users", uc.Get())
	eAuth.GET("users/:id", uc.GetById())
	eAuth.PUT("users/:id", uc.Update())
	eAuth.DELETE("users/:id", uc.Delete())

	// middleware.JWT([]byte(config.JWT_SECRET))
	eTask := e.Group("todo/")
	eTask.POST("task/register", tc.TaskRegister())
	eTask.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eTask.GET("task", tc.Get())
	eTask.GET("task/:id", tc.GetById())
	eTask.PUT("task/:id", tc.Update())
	eTask.DELETE("task/:id", tc.Delete())

	eProject := e.Group("")
	eProject.POST("projects/register", pc.ProjectRegister())
	eProject.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eProject.GET("projects", pc.Get())
	eProject.GET("projects/:id", pc.GetById())
	eProject.PUT("projects/:id", pc.Update())
	eProject.DELETE("projects/:id", pc.Delete())
	// e.POST("task/:id/completed", pc.TaskCompleted())
	// e.POST("task/:id/reopen", pc.TaskReopen())

}
