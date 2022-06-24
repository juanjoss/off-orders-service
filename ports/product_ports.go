package ports

import (
	"net/http"

	"github.com/juanjoss/off-orders-service/model"
)

/*
	Interfaces
*/

type ProductRepository interface {
	GetAllProducts() (GetAllProductsResponse, error)
	GetRandomProductFromUserSsd() (GetRandomProductFromUserSsdResponse, error)
	GetRandomProduct() (GetRandomProductResponse, error)
	CreateProductOrder(CreateProductOrderRequest) error
}

type ProductService interface {
	GetAllProducts() (GetAllProductsResponse, error)
	GetRandomProductFromUserSsd() (GetRandomProductFromUserSsdResponse, error)
	GetRandomProduct() (GetRandomProductResponse, error)
	CreateProductOrder(CreateProductOrderRequest) error
}

type ProductHandlers interface {
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetRandomProductFromUserSsd(w http.ResponseWriter, r *http.Request)
	GetRandomProduct(w http.ResponseWriter, r *http.Request)
	CreateProductOrder(w http.ResponseWriter, r *http.Request)
}

/*
	Service Models
*/

type GetAllProductsResponse struct {
	Products []*model.Product `json:"products"`
}

type GetRandomProductFromUserSsdResponse struct {
	SsdId   int    `json:"ssd_id"`
	Barcode string `json:"barcode"`
}

type GetRandomProductResponse model.Product

type CreateProductOrderRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
}
