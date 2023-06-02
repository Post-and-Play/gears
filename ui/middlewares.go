package ui

import (
	"github.com/gin-gonic/gin"
)

func AddDefaultHandlers(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
}
