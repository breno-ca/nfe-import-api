package main

import (
	"desafio-tecnico-backend/internal/config"
	"desafio-tecnico-backend/pkg/database"
	"desafio-tecnico-backend/pkg/routes"
	"desafio-tecnico-backend/pkg/security"
	"desafio-tecnico-backend/pkg/server"
	"desafio-tecnico-backend/pkg/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	default_config := &config.Config{}

	file_config, err := os.Open("./env.json")
	if err != nil {
		log.Panicln(err.Error())
	}

	jsonByte, err := io.ReadAll(file_config)
	if err != nil {
		log.Panicln(err.Error())
	}

	if err := json.Unmarshal(jsonByte, &default_config); err != nil {
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	conf := config.NewConfig(default_config)

	database_pool := database.NewBD(conf)
	if database_pool != nil {
		log.Print("Successfully connected")
	}

	service := service.NewAPIService(database_pool)

	err = security.SecretConfig(conf)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(conf)

	router := routes.ConfigRoutes(srv.SERVER, service)

	server.Run(router, srv, service)

	data, _ := json.Marshal(conf)

	fmt.Println(string(data))

}
