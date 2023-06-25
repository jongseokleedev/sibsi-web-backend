package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/controllers"
)

func ReceiverRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")

	//index.GET("/receivers", controllers.GetReceiver)
	index.GET("/receiver/:index", controllers.GetReceiver)
	index.POST("/receiver/new", controllers.CreateNewReceiver)
}
