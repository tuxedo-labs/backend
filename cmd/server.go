package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()
	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)
}
