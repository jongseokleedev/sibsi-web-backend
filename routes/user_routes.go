package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/controllers"
)

func UserRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")

	//index.GET("/receivers", controllers.GetReceiver)
	index.POST("/users", controllers.SignUp)
	index.POST("/users/token", controllers.SignIn)
	//index.POST("/receiver/new", controllers.CreateNewReceiver)
}
