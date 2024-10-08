package main

import (
	"fmt"
	"log"
	"os"
	"tuxedo/database"
	"tuxedo/routes"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
  fmt.Println("test")
	app := fiber.New(fiber.Config{
		AppName:      "Tuxedo BackEnd",
		ServerHeader: "Tuxedo",
		BodyLimit:    10 * 1024 * 1024,
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.Connect()

	routes.AutoMigrate()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:5173, https://tuxedo-frontend.vercel.app", // Mengizinkan semua asal
	// 	AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD",              // Mengizinkan semua metode
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization, x-token",      // Header yang diizinkan
	// 	ExposeHeaders:    "Content-Length",                                            // Header yang dapat diekspos
	// 	AllowCredentials: true,                                                        // Mengizinkan kredensial
	// }))

	app.Static("/", "./public")
	routes.SetupRouter(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
