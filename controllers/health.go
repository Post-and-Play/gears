package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Readiness(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
