package main

import (
	"goDemoApi/app/controllers"
	"goDemoApi/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load .env file")
	}

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/users/:id", controllers.UsersShow)
	v1.Post("/contact-requests", controllers.ContactRequestsStore)

	app.Listen(":3000")
}
