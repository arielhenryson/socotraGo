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
		routePath := route.Path
		r.GET(routePath, func(c *gin.Context) {
			c.String(http.StatusOK, routePath)
		})
	}


	return r
}