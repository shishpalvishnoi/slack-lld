package http_server

import "github.com/gin-gonic/gin"

func healthCheck(c *gin.Context) {
	c.String(200, "OK")
}
