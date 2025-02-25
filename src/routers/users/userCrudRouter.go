package routers

import (
	handlers "base-go-app/src/apps/users/handlers/user"
	"base-go-app/src/apps/users/models"
	middlewares "base-go-app/src/middlewares/authentication"

	"github.com/gin-gonic/gin"
)

func UserCrudRouter(router *gin.RouterGroup) {
	// User Crud
	userHandler := handlers.CreateUserCrudHandler()
	allowedRoles := []string{
		models.Superuser, models.Manager,
	}
	{
		usersRouter := router.Group("/users")
		usersRouter.Use(
			middlewares.AuthMiddleware(),
			middlewares.CheckRoleMiddleware(allowedRoles...),
		)
		{
			usersRouter.GET(
				"/",
				userHandler.ListAll,
			)
		}
	}
}
