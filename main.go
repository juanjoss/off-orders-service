package main

import (
	"os"

	"github.com/juanjoss/off-orders-service/handlers"
	"github.com/juanjoss/off-orders-service/repository"
	"github.com/juanjoss/off-orders-service/server"
	"github.com/juanjoss/off-orders-service/services"
)

func main() {
	pr := repository.NewProductRepository()
	ps := services.NewProductService(pr)
	ph := handlers.NewProductHandlers(ps)
	server.NewServer(ph).ListenAndServe(":" + os.Getenv("SERVICE_PORT"))
}
