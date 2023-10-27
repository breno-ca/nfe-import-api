package main

import (
	"desafio-tecnico-backend/internal/config"
	"desafio-tecnico-backend/pkg/database"
	"desafio-tecnico-backend/pkg/routes"
	"desafio-tecnico-backend/pkg/security"
	"desafio-tecnico-backend/pkg/server"
	"desafio-tecnico-backend/pkg/service"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {

	default_config := &config.Config{}

	// Abre o arquivo env
	file_config, err := os.Open("./env.json")
	if err != nil {
		log.Panicln(err.Error())
	}

	// Realiza a Leitura do Arquivo env
	jsonByte, err := io.ReadAll(file_config)
	if err != nil {
		log.Panicln(err.Error())
	}

	// Atribui as configurações do arquivo a variável
	if err := json.Unmarshal(jsonByte, &default_config); err != nil {
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	// Aplica as configurações
	conf := config.NewConfig(default_config)

	database_pool := database.NewBD(conf)
	if database_pool != nil {
		log.Print("Successfully connected")
	}

	// Cria um novo serviço
	service := service.NewAPIService(database_pool)

	// Configura o segredo JWT
	err = security.SecretConfig(conf)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(conf)                      // Cria o novo server com as configurações
	router := routes.ConfigRoutes(srv.SERVER, service) // Configura as rotas do arquivo de rotas
	server.Run(router, srv, service)                   // Inicia o servidor

}
