package main

import (
	"github.com/arielhenryson/socotraGo/config"
	"github.com/arielhenryson/socotraGo/core"
	"github.com/arielhenryson/socotraGo/routes"
)

func main() {
	core.ServerInit(config.AppConfig, routes.Routes)
}
