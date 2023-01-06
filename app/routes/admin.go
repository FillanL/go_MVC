package routes

import (
	"github.com/FillanL/creatturlinks/app/controller"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutesSetup(route fiber.Router){
	controller := controller.Init()

	route.Get("/", controller.AdminController.Get)

}