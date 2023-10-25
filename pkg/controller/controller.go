package controller

import (
	"desafio-tecnico-backend/pkg/entity"
	"desafio-tecnico-backend/pkg/security"
	"desafio-tecnico-backend/pkg/service"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func CreateProduto(c *gin.Context, service service.ProdutoServiceInterface) {
	var produto *entity.Prod
	err := c.ShouldBind(&produto)
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}
	ctx := c.Request.Context()
	err = service.CreateProduto(produto, ctx)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	send(c, http.StatusCreated, gin.H{
		"message": "Produto cadastrado com sucesso",
	})
}

func ImportNFe(c *gin.Context, service service.NFeImportServiceInterface) {
	xmlData, err := c.GetRawData()
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	var nfe entity.NFe
	if err := xml.Unmarshal(xmlData, &nfe); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	err = service.ImportNFeXML(&nfe, ctx)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	send(c, http.StatusOK, gin.H{
		"message": "XML importado com sucesso",
	})
}
