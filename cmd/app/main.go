package main

import (
	"base-go-app/config/settings/environment"
	"base-go-app/src/database"
	"base-go-app/src/routers"

	"github.com/gin-gonic/gin"

	securityMiddlewares "base-go-app/src/middlewares/security"

	_ "base-go-app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// For Swagger UI Authentication
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.
func main() {
	// Initialize Environment Variables
	environment.InitalizeDotEnv()
	// Connect to Postgres Database
	database.ConnectPostgres()
	// Setup Router and Run
	engine := setupRouter()
	// Swagger Docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run the server
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
