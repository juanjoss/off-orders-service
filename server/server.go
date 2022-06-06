package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanjoss/off-orders-service/ports"
)

const apiPrefix = "/api/"

type Server struct {
	productHandlers ports.ProductHandlers
	userHandlers    ports.UserHandlers
	router          *mux.Router
}

func NewServer(ph ports.ProductHandlers, uh ports.UserHandlers) *Server {
	return &Server{
		productHandlers: ph,
		userHandlers:    uh,
		router:          mux.NewRouter().PathPrefix(apiPrefix).Subrouter(),
	}
}

func (s *Server) RegisterRoutes() {
	// users
	s.router.HandleFunc("/register", s.userHandlers.Register).Methods(http.MethodPost)
	s.router.HandleFunc("/users/ssds/products", s.userHandlers.AddProductToSSD).Methods(http.MethodPost)
	s.router.HandleFunc("/users/ssds/random", s.userHandlers.RandomSSD).Methods(http.MethodGet)

	// products
	s.router.HandleFunc("/products", s.productHandlers.GetAll).Methods(http.MethodGet)
	s.router.HandleFunc("/products/randomProductFromUserSSD", s.productHandlers.GetRandomProductFromUserSsd).Methods(http.MethodGet)
	s.router.HandleFunc("/products/random", s.productHandlers.Random).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe(addr string) {
	s.RegisterRoutes()

	log.Println("Starting server at", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatal(err.Error())
	}
}
