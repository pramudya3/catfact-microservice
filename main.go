package main

import (
	"go-microservice/internal/handler/api"
	"go-microservice/internal/logging"
	"go-microservice/internal/service"
	"log"
)

//microservice intercomunication each other like grpc, proto buffer,

func main() {
	svc := service.NewCatFactService("https://catfact.ninja/fact")
	svc = logging.NewCatFactLogging(svc)

	apiServer := api.NewApiServer(svc)
	log.Fatal(apiServer.Start(":3000"))
}
