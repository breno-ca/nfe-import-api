package routes

import (
	"desafio-tecnico-backend/pkg/controller"
	"desafio-tecnico-backend/pkg/middleware"
	"desafio-tecnico-backend/pkg/service"

	"github.com/gin-gonic/gin"
)

// Rotas do sistema
func ConfigRoutes(router *gin.Engine, service service.APIServiceInterface) *gin.Engine {

	router.Use(middleware.CORS())

	main := router.Group("dtb")
	{

		userGroup := main.Group("/v1/user")
		{
			userGroup.POST("/login", func(c *gin.Context) {
				controller.Login(c, service)
			})
		}

		nfeGroup := main.Group("/v1/nfe")
		{
			nfeGroup.POST("/importar", middleware.Auth(), func(c *gin.Context) {
				controller.ImportNFeXML(c, service)
			})
		}

	}
	return router
}
