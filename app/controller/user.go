package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct{

}
func (uh *UserHandlers) GetUser(c *fiber.Ctx)error{return errors.New("something happened")}
func (uh *UserHandlers) CreateUserFromProvider(c *fiber.Ctx)error{
	return errors.New("something happened")
}

func (uh *UserHandlers) CreateUser(c *fiber.Ctx)error{
	return errors.New("something happened")
}