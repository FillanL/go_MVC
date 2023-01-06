package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type LinkHandlers struct{
}

func (linkController *LinkHandlers) Get(c *fiber.Ctx)error{
	return errors.New("something happened")
}
func (linkController *LinkHandlers) PostLink(c *fiber.Ctx)error{
	return errors.New("something happened")
}
func (linkController *LinkHandlers) GetLinks(c *fiber.Ctx)error{
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "ok",
		},
	)
	return errors.New("something happened")
}
