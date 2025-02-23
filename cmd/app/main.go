package main

import (
	"base-go-app/src/apps/users/models"
	"base-go-app/src/common/utils/environment"
	"base-go-app/src/database"

	"github.com/gin-gonic/gin"

	jwtHandlers "base-go-app/src/apps/users/handlers/jwt"
)

func main() {
	environment.InitalizeEnv()
	database.ConnectPostgres()

	r := gin.Default()
	// Api Group
	api := r.Group("/api")
	// Authentication
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", jwtHandlers.LoginHandler)
			auth.POST("/refresh", jwtHandlers.RefreshHandler)
		}
	}

	usersGroup := api.Group("/users")
	{
		{
			usersGroup.GET("/", func(c *gin.Context) {
				var users []models.User
				database.PostgresContext.Find(&users)
				c.JSON(200, gin.H{"data": users})
			})
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}
