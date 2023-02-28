package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary      Healthcheck
// @Description  Route to health check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /healthz [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Readiness godoc
// @Summary      Healthcheck
// @Description  Route to readiness check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /readiness [get]
func Readiness(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
