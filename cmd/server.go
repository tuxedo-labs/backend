package main

import (
	"os"
	"tuxedo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()
	routes.AutoMigrate()
	routes.SetupRouter(app)
	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)
}
