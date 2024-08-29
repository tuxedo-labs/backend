package middleware

import (
	"net/http"
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

	_, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("usersInfo", claims)
	c.Locals("role", claims["role"])

	return c.Next()
}
