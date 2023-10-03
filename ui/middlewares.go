package ui

import (
	"github.com/gin-gonic/gin"
)

func AddDefaultHandlers(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,PUT,OPTIONS,POST,DELETE,PATCH")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")

	if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(204)
        return
    }

	c.Next()
}
