package utils

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Verify(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	validToken := os.Getenv("ACCESS_TOKEN")
	if token != validToken {
		return false
	}
	return true
}
