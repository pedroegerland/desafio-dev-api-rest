package service

import (
	"banktest_signin/src/entity"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var credential entity.Credentials

	if err := c.BodyParser(&credential); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if entity.Users[credential.Cpf] != "" {
		fmt.Println("User already exists!")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.Response{Msg: "User already exists!"})
	}

	entity.Users[credential.Cpf] = credential.Password

	fmt.Println("User created!")

	return c.Status(fiber.StatusOK).JSON(entity.Response{Msg: "User created!"})
}
