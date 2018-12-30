package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Routes = []Route{
	{
		Path: "/",
	},
	{
		Path: "/page2",
	},
	{
		Path: "/page3",
	},
}


func (p Route) Controller(c *gin.Context) {
	c.String(http.StatusOK, "test")
}