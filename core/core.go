package core

import (
	"fmt"
	appConfig "github.com/arielhenryson/socotraGo/config"
	"strconv"
)


func ServerInit(config appConfig.Config) {
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	fmt.Println("Starting " + config.AppName)
	r.Run(":" + strconv.Itoa(config.HttpPort))
}