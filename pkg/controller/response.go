package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"status":     status,
		"statusText": http.StatusText(status),
		"error":      err.Error(),
		"path":       c.FullPath(),
	})
}

func send(c *gin.Context, code int, obj any) {
	c.JSON(code, obj)
}

// func sendNoContent(c *gin.Context) {
// 	c.Status(http.StatusNoContent)
// }
