package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/controllers"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/auth"
)

func GiftRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")
	index.Use(auth.TokenAuthMiddleware)
	{
		index.GET("/gift/:id", controllers.GetGift)
		index.GET("/gift/:id/providers", controllers.GetAllProviders)
		index.POST("/gift/new", controllers.CreateNewGift)
		index.POST("/gift/:id", controllers.UpdateGift)
		index.POST("/gift/:id/provider", controllers.CreateNewProvider)
		index.POST("/gift/:id/provider/:providerID", controllers.RemoveProvider)
	}
}
