package routes

import (
	routes "github.com/FillanL/creatturlinks/app/routes/auth"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App)*fiber.App{
	api := app.Group("/api/v1")
	
	linkRoutes := api.Group("/link")
	LinkRoutesSetup(linkRoutes)

	authRoutes := api.Group("/auth")
	routes.AuthRoutesSetup(authRoutes)
	
	// meRoutes := api.Group("/me")
	// meRoutes.Use()
	admin := api.Group("/admin")
	AdminRoutesSetup(admin)

	return app
}