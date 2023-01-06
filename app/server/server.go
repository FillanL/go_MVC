package server

import (
	"github.com/FillanL/creatturlinks/app/config"
	"github.com/FillanL/creatturlinks/app/database"
	"github.com/FillanL/creatturlinks/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() *fiber.App {
	app := fiber.New()

	app.Use(
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders:  "Origin, Content-Type, Accept",
		}),
		limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}),
	)
	config.LoadEnvironmentFile()
	database.SetUpDatabase()
	routes.SetRoutes(app)
	
	return app
}