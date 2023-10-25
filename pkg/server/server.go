package server

import (
	"desafio-tecnico-backend/internal/config"
	"desafio-tecnico-backend/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	SRV_PORT string
	SERVER   *gin.Engine
}

func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

func Run(router *gin.Engine, server Server, service service.UserServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)
	log.Fatal(router.Run(":" + server.SRV_PORT))
}
