package routes

import (
	"github.com/FillanL/creatturlinks/app/controller"
	"github.com/FillanL/creatturlinks/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func LinkRoutesSetup(route fiber.Router){
	controller := controller.Init()

	route.Get("/", controller.LinkController.GetLinks)
	route.Post("/", middleware.IsAuthenticated ,controller.LinkController.PostLink)
	route.Post("/user", controller.LinkController.GetLinks)

}