package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers struct{
}
 func (ah AuthHandlers) IsAuthenticated(c *fiber.Ctx)error{return errors.New("something happened")}