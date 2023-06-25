package main

import (
	"fmt"

	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"github.com/jongseokleedev/sibsi-web-backend/server/routes"
)

func main() {
	fmt.Println("Sibsi API Server Start ........")
	//configs.ConnectDB()
	//fmt.Println("connect DB Done ...")
	routes.InitRoutes()
	routes.Run(configs.Port)
}
