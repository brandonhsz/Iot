package main

import (
	"main/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	loadEnv()
	router.SetupRouter(app)

	app.Listen(os.Getenv("PORT"))

}

func loadEnv() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
