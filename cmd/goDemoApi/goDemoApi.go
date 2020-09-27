package main

import (
	"goDemoApi/database"
	"goDemoApi/internal/goDemoApi/http"
	"goDemoApi/queue"
	"log"

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

	http.Init()
	http.Routes()
	http.Serve()
}
