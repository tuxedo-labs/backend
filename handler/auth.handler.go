package handler

import (
	"fmt"
	"tuxedo/models/request"
	"tuxedo/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	if errValidate := services.ValidateLogin(loginRequest); errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	user, err := services.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	token, errGenerateToken := services.GenerateJWTToken(user)
	if errGenerateToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Wrong credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	registerRequest := new(request.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return err
	}

	if errValidate := services.ValidateRegister(registerRequest); errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	if err := services.HashAndStoreUser(registerRequest); err != nil {
		if err.Error() == fmt.Sprintf("user with email %s already exists", registerRequest.Email) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email already in use",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success register",
	})
}