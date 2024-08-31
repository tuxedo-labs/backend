package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"tuxedo/services"
)

func GetProfile(c *fiber.Ctx) error {
	usersInfo := c.Locals("usersInfo")
	if usersInfo == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := usersInfo.(jwt.MapClaims)

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "User ID not found",
		})
	}
	userId := uint(idFloat)

	data, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching user data",
		})
	}

	profile, err := services.BuildUserProfile(data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error building user profile",
		})
	}

	return c.JSON(fiber.Map{
		"data": profile,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	return nil
}
