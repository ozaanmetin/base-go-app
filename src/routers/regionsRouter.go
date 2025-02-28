package routers

import (
	handlers "base-go-app/src/apps/regions/handlers/country"
	"base-go-app/src/apps/users/models"
	middlewares "base-go-app/src/middlewares/authentication"

	"github.com/gin-gonic/gin"
)

func RegionsRouter(router *gin.RouterGroup) {
	countryHandler := handlers.CreateCountryHandler()
	{
		allowedRoles := []string{
			models.Superuser,
			models.Manager,
		}
		regionsRouter := router.Group("/regions")
		regionsRouter.Use(
			middlewares.AuthMiddleware(),
			middlewares.CheckRoleMiddleware(allowedRoles...),
		)
		{
			regionsRouter.GET("/country", countryHandler.ListAll)
			regionsRouter.GET("/country/:id", countryHandler.Retrieve)
		}

	}
}
