package main

import (
	"Project-REST-API/configs"
	"Project-REST-API/delivery/controllers/routes"
	uc "Project-REST-API/delivery/controllers/user"
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

	userController := uc.New(userRepo)

	e := echo.New()

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
