package main

import (
	"Project-REST-API/configs"
	ac "Project-REST-API/delivery/controllers/auth"
	pc "Project-REST-API/delivery/controllers/project"
	"Project-REST-API/delivery/controllers/routes"
	tc "Project-REST-API/delivery/controllers/task"
	uc "Project-REST-API/delivery/controllers/user"
	authRepo "Project-REST-API/repository/auth"
	projectRepo "Project-REST-API/repository/project"
	taskRepo "Project-REST-API/repository/task"
	userRepo "Project-REST-API/repository/user"
	"Project-REST-API/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	taskRepo := taskRepo.New(db)
	projectRepo := projectRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	taskController := tc.New(taskRepo)
	projectController := pc.New(projectRepo)

	e := echo.New()

	routes.RegisterPath(e, authController, userController, taskController, projectController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
