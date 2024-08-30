package handler

import (
	"fmt"
	"net/http"
	"tuxedo/models/request"
	"tuxedo/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	usersInfo := c.Locals("usersInfo")
	if usersInfo == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	fmt.Printf("UsersInfo: %+v\n", usersInfo)

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

	profile := request.UserProfile{
		Name:      data.Name,
		Email:     data.Email,
		Role:      data.Role,
		CreatedAt: data.CreatedAt.Format("2006-01-01"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-01"),
	}

	return c.JSON(fiber.Map{
		"data": profile,
	})
}
