package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juanjoss/off-orders-service/ports"
)

const apiPrefix = "/api/products"

type Server struct {
	productHandlers ports.ProductHandlers
	router          *mux.Router
	port            string
}

func NewServer(ph ports.ProductHandlers) *Server {
	return &Server{
		productHandlers: ph,
		router:          mux.NewRouter().PathPrefix(apiPrefix).Subrouter(),
		port:            ":" + os.Getenv("SERVICE_PORT"),
	}
}

func (s *Server) RegisterRoutes() {
	// products
	s.router.HandleFunc("", s.productHandlers.GetAllProducts).Methods(http.MethodGet)
	s.router.HandleFunc("/randomProductFromUserSSD", s.productHandlers.GetRandomProductFromUserSsd).Methods(http.MethodGet)
	s.router.HandleFunc("/random", s.productHandlers.GetRandomProduct).Methods(http.MethodGet)

	// orders
	s.router.HandleFunc("/orders", s.productHandlers.CreateProductOrder).Methods(http.MethodPost)
}

func (s *Server) ListenAndServe() {
	s.RegisterRoutes()

	log.Println("Starting server at", s.port)
	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Fatal(err.Error())
	}
}
