package utils

import (
	"github.com/gin-gonic/gin"
)

func SaveNodeRequestIp(c *gin.Context) {
	node := c.GetHeader("node-ip")
	if node != "" {
		AppendNodesFromIPString(node)
	}
	c.Next()
}
