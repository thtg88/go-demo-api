package main

import (
	"flag"
	"fmt"
	"goDemoApi/app/controllers"
	"goDemoApi/database"
	"goDemoApi/queue"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var prod = flag.Bool("prod", false, "Enable prefork in Production")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load .env file")
	}

	database.Connect()
	database.AutoMigrate()

	queue.InitializeRedis()

	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
		},
	}))

	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/users/:id", controllers.UsersShow)
	v1.Post("/contact-requests", controllers.ContactRequestsStore)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(fmt.Sprintf(":%s", port))
}
