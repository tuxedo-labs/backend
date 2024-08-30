package routes

import (
	"tuxedo/config"
	"tuxedo/handler"
	"tuxedo/middleware"
	"tuxedo/models/entity"

	"github.com/gofiber/fiber/v2"
)

var auth = middleware.Auth
var admin = middleware.AdminRole

func SetupRouter(r *fiber.App) {
	app := r.Group("/api")
	// authentication
	app.Post("/auth/login", handler.Login)
	app.Post("/auth/register", handler.Register)

	//users
	app.Get("/users/profile", auth, handler.GetProfile)
}

func AutoMigrate() {
	config.RunMigrate(&entity.Users{})
}
