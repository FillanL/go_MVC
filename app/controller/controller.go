package controller

import "github.com/gofiber/fiber/v2"

type Controllers struct{
	LinkController LinkHandlers
	UserContoller UserHandlers
	GoogleController GoogleHandlers
	AdminController AdminHandle
}
func Init()Controllers{
	return Controllers{}

}
func gernerateJsonError(message string, c *fiber.Ctx, httpStatus int) error{
	return c.Status(httpStatus).JSON(
		fiber.Map{
			"message": message,
		},
	)
}