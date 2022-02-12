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

	//=========================================================
	//ROUT USERS
	e.POST("users/register", uc.UserRegister())
	e.POST("users/login", uc.Login())
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("users", uc.Get())
	eAuth.GET("users/:id", uc.GetById())
	eAuth.PUT("users/:id", uc.Update())
	eAuth.DELETE("users/:id", uc.Delete())

	//===========================================================
	//ROUTE TASK
	eTask := e.Group("todo/")
	eTask.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eTask.POST("task/register", tc.TaskRegister())
	eTask.GET("task", tc.Get())
	eTask.GET("task/:id", tc.GetById())
	eTask.PUT("task/:id", tc.Update())
	eTask.DELETE("task/:id", tc.Delete())
	// e.POST("task/:id/completed", tc.TaskCompleted())
	// e.POST("task/:id/reopen", tc.TaskReopen())

	//===========================================================
	//ROUTE PROJECT
	e.POST("projects/register", pc.ProjectRegister())
	eProject := e.Group("")
	eProject.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eProject.GET("projects", pc.Get())
	eProject.GET("projects/:id", pc.GetById())
	eProject.PUT("projects/:id", pc.Update())
	eProject.DELETE("projects/:id", pc.Delete())

}
