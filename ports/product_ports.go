package ports

import (
	"net/http"

	"github.com/juanjoss/off-orders-service/model"
)

type ProductRepository interface {
	GetAll() ([]*model.Product, error)
	GetRandomProductFromUserSsd() (int, int, string, error)
	Random() (*model.Product, error)
}

type ProductService interface {
	GetAll() ([]*model.Product, error)
	GetRandomProductFromUserSsd() (int, int, string, error)
	Random() (*model.Product, error)
}

type ProductHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetRandomProductFromUserSsd(w http.ResponseWriter, r *http.Request)
	Random(w http.ResponseWriter, r *http.Request)
}
