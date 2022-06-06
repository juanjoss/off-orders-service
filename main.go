package main

import (
	"os"

	"github.com/juanjoss/off-orders-service/handlers"
	"github.com/juanjoss/off-orders-service/repository"
	"github.com/juanjoss/off-orders-service/server"
	"github.com/juanjoss/off-orders-service/services"
)

func main() {
	// repositories
	pr := repository.NewProductRepository()
	ur := repository.NewUserRepository()

	// services
	ps := services.NewProductService(pr)
	us := services.NewUserService(ur)

	// handlers
	ph := handlers.NewProductHandlers(ps)
	uh := handlers.NewUserHandlers(us)

	// server
	server.NewServer(ph, uh).ListenAndServe(":" + os.Getenv("SERVICE_PORT"))
}
