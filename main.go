package main

import (
	"github.com/arielhenryson/socotraGo/config"
	"github.com/arielhenryson/socotraGo/core"
)

func main() {
	core.ServerInit(config.InitConfig())
}
