package routes

import (
	"github.com/GCI-js/bangalzoo/server/controllers"
	"github.com/gin-gonic/gin"
)

func ReceiverRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")

	//index.GET("/receivers", controllers.GetReceiver)
	index.GET("/receiver/:index", controllers.GetReceiver)
}
