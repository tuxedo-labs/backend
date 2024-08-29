package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AdminRole(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role == "member" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return c.Next()
}
