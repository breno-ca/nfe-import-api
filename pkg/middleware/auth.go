package middleware

import (
	"desafio-tecnico-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware de autenticação de acessos das rotas
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token de autenticacao ausente"})
			c.Abort()
			return
		}

		token := header[len(bearerSchema):]

		err := security.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token de autenticacao invalido"})
			c.Abort()
			return
		}
	}
}
