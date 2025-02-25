package routers

import (
	"github.com/gin-gonic/gin"
	"base-go-app/src/apps/users/handlers/jwt"
)


func AuthenticationRouter(router *gin.RouterGroup) {
	// Authentication
	{
		auth := router.Group("/auth")
		{
			auth.POST("/login", handlers.LoginHandler)
			auth.POST("/refresh", handlers.RefreshHandler)
		}
	}
}
