package routers

import (
	handlers "base-go-app/src/apps/users/handlers/user"
	"base-go-app/src/apps/users/models"
	middlewares "base-go-app/src/middlewares/authentication"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	// User Crud
	userHandler := handlers.CreateUserCrudHandler()
	{
		allowedRoles := []string{
			models.Superuser,
			models.Manager,
		}
		usersRouter := router.Group("/users")
		usersRouter.Use(
			middlewares.AuthMiddleware(),
			middlewares.CheckRoleMiddleware(allowedRoles...),
		)
		{
			usersRouter.GET("/", userHandler.ListAll)
			usersRouter.GET("/:id", userHandler.Retrieve)
			usersRouter.POST("/", userHandler.Create)
			usersRouter.PATCH("/:id", userHandler.Update)
			usersRouter.DELETE("/:id", userHandler.Delete)
		}
	}
}
