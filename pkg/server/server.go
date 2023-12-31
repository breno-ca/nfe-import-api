package server

import (
	"desafio-tecnico-backend/internal/config"
	"desafio-tecnico-backend/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
)

// Struct com as configurações do servidor
type Server struct {
	SRV_PORT string
	SERVER   *gin.Engine
}

// Construtor do Server
func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

// Inicia o router Gin-gonic
func Run(router *gin.Engine, server Server, service service.UserServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)
	log.Fatal(router.Run(":" + server.SRV_PORT))
}
