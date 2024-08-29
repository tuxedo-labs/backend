package main

import (
	"log"
	"os"
	"tuxedo/database"
	"tuxedo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()
	routes.AutoMigrate()
	routes.SetupRouter(app)
	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)
}
