package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura de mensagem para as requisições
func sendError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"status":     status,
		"statusText": http.StatusText(status),
		"error":      err.Error(),
		"path":       c.FullPath(),
	})
}

// Envio de mensagem para as requisições
func send(c *gin.Context, code int, obj any) {
	c.JSON(code, obj)
}
