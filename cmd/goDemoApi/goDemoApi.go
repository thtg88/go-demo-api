package main

import (
	"fmt"
	"goDemoApi/app/controllers"
	"goDemoApi/database"
	"goDemoApi/queue"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load .env file")
	}

	database.Connect()
	database.AutoMigrate()

	queue.InitializeRedis()

	app := gin.Default()

	authorized := app.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	}))

	api := authorized.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/users/:id", controllers.UsersShow)
	v1.POST("/contact-requests", controllers.ContactRequestsStore)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Run(fmt.Sprintf(":%s", port))
}
