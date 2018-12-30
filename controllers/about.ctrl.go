package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func AboutCtrl(c *gin.Context) {
	c.String(http.StatusOK, "About page")
}