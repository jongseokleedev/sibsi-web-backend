package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/controllers"
)

func UserRoutes(rg *gin.RouterGroup) {
	index := rg.Group("")

	index.POST("/users", controllers.SignUp)
	index.POST("/users/login", controllers.SignIn)
	index.POST("/users/logout", controllers.Logout)
}
