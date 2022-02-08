package main

import (
	"Project-REST-API/configs"
	"Project-REST-API/delivery/controllers/routes"
	tc "Project-REST-API/delivery/controllers/task"
	uc "Project-REST-API/delivery/controllers/user"
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

	userRepo := userRepo.New(db)
	taskRepo := taskRepo.New(db)

	userController := uc.New(userRepo)
	taskController := tc.New(taskRepo)

	e := echo.New()

	routes.RegisterPath(e, userController, taskController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
