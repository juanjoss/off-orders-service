package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanjoss/off-orders-service/ports"
)

const apiPrefix = "/api/products"

type Server struct {
	productHandlers ports.ProductHandlers
	router          *mux.Router
}

func NewServer(ph ports.ProductHandlers) *Server {
	return &Server{
		productHandlers: ph,
		router:          mux.NewRouter().PathPrefix(apiPrefix).Subrouter(),
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

func (s *Server) ListenAndServe(addr string) {
	s.RegisterRoutes()

	log.Println("Starting server at", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatal(err.Error())
	}
}
