package main

import (
	"log"

	"github.com/FillanL/creatturlinks/app/server"
)

// func test(c *fiber.Ctx) error{
// 	fmt.Println("test")
// 	return c.Next()
// }

func main(){
	// app := fiber.New()

	// app.Use(
	// 	logger.New(logger.Config{
	// 		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// 	}),
	// 	cors.New(cors.Config{
	// 		AllowOrigins: "*",
	// 		AllowHeaders:  "Origin, Content-Type, Accept",
	// 	}),
	// 	limiter.New(limiter.Config{
	// 	Max: 100,
	// 	LimitReached: func(c *fiber.Ctx) error {
	// 		return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
	// 			"status":  "fail",
	// 			"message": "You have requested too many in a single time-frame! Please wait another minute!",
	// 		})
	// 	},
	// }),
	// )
	// config.LoadEnvironmentFile()
	// database.SetUpDatabase()
	// routes.SetRoutes(app)
	app := server.Setup()
	
	log.Fatal(app.Listen(":8080"))
}
