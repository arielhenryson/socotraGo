package core

import (
	"github.com/arielhenryson/socotraGo/routes"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()


	routes.SetupAppRouter(r)



	return r
}