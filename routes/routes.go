package routes

import (
	"github.com/arielhenryson/socotraGo/controllers"
	"github.com/arielhenryson/socotraGo/middlewares"
	"github.com/gin-gonic/gin"
)


func SetupAppRouter(r *gin.Engine)  {
	r.GET("/", controllers.IndexCtrl)
	r.GET("/about", middlewares.TestMid(), controllers.AboutCtrl)
}