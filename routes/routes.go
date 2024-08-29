package routes

import (
	"tuxedo/config"
	"tuxedo/handler"
	"tuxedo/models/entity"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	app := r.Group("/api")
	// authentication
	app.Post("/auth/login", handler.Login)
	app.Post("/auth/register", handler.Register)
}

func AutoMigrate() {
	config.RunMigrate(&entity.Users{})
}
