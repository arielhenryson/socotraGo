package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexCtrl(c *gin.Context) {
	c.String(http.StatusOK, "Home page")
}