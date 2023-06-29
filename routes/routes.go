package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
)

var router = gin.Default()

func Run(port uint64) {
	router.Run(":" + strconv.FormatUint(port, 10))
}

func InitRoutes() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}                              // allowed methods
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"} // allowed headers
	router.Use(cors.New(config))

	docs.SwaggerInfo.BasePath = "/sibsi"
	docs.SwaggerInfo.Title = "sibsi API"
	v1 := router.Group("sibsi")
	router.GET("/", getVersion)
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		ReceiverRoutes(v1)
		GiftRoutes(v1)
		UserRoutes(v1)
	}

}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "v1.0", "service": "Sibsi API Server"})
}

func SetMode(mode string) {
	gin.SetMode(mode)
}
