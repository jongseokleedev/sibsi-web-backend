package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/controllers"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/auth"
)

func ReceiverRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")
	index.Use(auth.TokenAuthMiddleware)
	{
		index.GET("/receiver/:index", controllers.GetReceiver)
		index.POST("/receiver/new", controllers.CreateNewReceiver)
	}
	//index.GET("/receivers", controllers.GetReceiver)
}
