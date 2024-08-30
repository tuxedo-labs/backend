package middleware

import (
	"net/http"
	"tuxedo/database"
	"tuxedo/models/entity"
	"tuxedo/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	userID := claims["id"].(uint)
	var user entity.Users
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	c.Locals("usersInfo", claims)
	c.Locals("role", claims["role"])
	return c.Next()
}
