package http

import (
	"fmt"
	"goDemoApi/internal/goDemoApi/http/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

// Instance return an instance of the HTTP app
func Instance() *gin.Engine {
	return app
}

// Init initializes the Gin default internal app
func Init() {
	app = gin.Default()
}

// Routes define the routes and middlewares for the existing app
func Routes() {
	authorized := app.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	}))

	api := authorized.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/users/:id", controllers.UsersShow)
	v1.POST("/contact-requests", controllers.ContactRequestsStore)
}

// Serve serves the Gin app from the port environment variable (if provided)
func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Run(fmt.Sprintf(":%s", port))
}
