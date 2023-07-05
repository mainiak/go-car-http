package internal

import (
	"github.com/gin-gonic/gin"
)

func serve_car(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test"})
}
