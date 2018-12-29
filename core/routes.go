package core

import (
	"github.com/arielhenryson/socotraGo/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SetupRouter(routes []routes.Route) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()


	// loop throw the app routes array
	for _, route := range routes {
		r.GET(route.Path, func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World")
		})
	}


	return r
}