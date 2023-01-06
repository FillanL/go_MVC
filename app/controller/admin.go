package controller

import (
	"github.com/FillanL/creatturlinks/app/models"
	"github.com/gofiber/fiber/v2"
)

type AdminHandle struct{}

func (ah *AdminHandle) Get(c *fiber.Ctx)error{
userModel := models.User{}
	// check if admin

	
	users:= userModel.GetUsers()
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message":"ok",
			"users": users,
		},
	)
}