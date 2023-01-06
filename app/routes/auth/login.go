package routes

import (
	"github.com/FillanL/creatturlinks/app/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutesSetup(route fiber.Router){
	google := route.Group("/google")
	googleRoutes(google)
}

func googleRoutes(route fiber.Router){
	controller := controller.Init()
	route.Get("/callback", controller.GoogleController.Callback)
	route.Get("/login", controller.GoogleController.Login)
}