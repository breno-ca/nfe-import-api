package controller

import (
	"desafio-tecnico-backend/pkg/entity"
	"desafio-tecnico-backend/pkg/security"
	"desafio-tecnico-backend/pkg/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller responsável pelo Login
func Login(c *gin.Context, service service.UserServiceInterface) {

	var user *entity.User

	err := c.ShouldBind(&user)
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	hash, err := service.Login(user, ctx)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	err = security.ValidatePassword(hash, user.Senha)
	if err != nil {
		sendError(c, http.StatusUnauthorized, errors.New("incorrect credentials"))
		return
	}

	token, err := security.NewToken(user.CNPJ)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	send(c, http.StatusOK, gin.H{
		"token": token,
	})

}

// Controller responsável pela importação da NFe
func ImportNFeXML(c *gin.Context, service service.NFeImportServiceInterface) {
	xmlData, err := c.GetRawData()
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	userCNPJ, _ := security.GetUserCNPJ(c)

	err = service.ImportNFeXML(xmlData, c, userCNPJ)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	send(c, http.StatusOK, gin.H{
		"message": "XML importado com sucesso",
	})
}
