package main

import (
	"fmt"

	"github.com/GCI-js/bangalzoo/server/configs"
	"github.com/GCI-js/bangalzoo/server/routes"
)

func main() {
	fmt.Println("Sibsi API Server Start ........")
	//configs.ConnectDB()
	//fmt.Println("connect DB Done ...")
	routes.InitRoutes()
	routes.Run(configs.Port)
}
