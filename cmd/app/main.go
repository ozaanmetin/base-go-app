package main

import (
	"base-go-app/config/settings/environment"
	"base-go-app/src/database"
	"base-go-app/src/routers"

	"github.com/gin-gonic/gin"

	securityMiddlewares "base-go-app/src/middlewares/security"
)

func main() {
	// Initialize Environment Variables
	environment.InitalizeDotEnv()
	// Connect to Postgres Database
	database.ConnectPostgres()
	// Setup Router and Run
	engine := setupRouter()
	engine.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Security Middlewares
	r.Use(securityMiddlewares.StrictHostValidationMiddleware())
	r.Use(securityMiddlewares.CorsMiddleware())
	// Api Group
	api := r.Group("/api")
	// Authentication
	routers.AuthenticationRouter(api)
	routers.UsersRouter(api)
	routers.RegionsRouter(api)
	return r
}
