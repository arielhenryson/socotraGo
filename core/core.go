package core

import (
	"fmt"
	appConfig "github.com/arielhenryson/socotraGo/config"
	"github.com/arielhenryson/socotraGo/routes"
	"strconv"
)


func ServerInit(config appConfig.Config, routes []routes.Route) {
	r := SetupRouter(routes)
	// Listen and Server in 0.0.0.0:8080
	fmt.Println("Starting " + config.AppName)
	r.Run(":" + strconv.Itoa(config.HttpPort))
}