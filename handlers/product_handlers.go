package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/juanjoss/off-orders-service/ports"
)

type ProductHandlers struct {
	productService ports.ProductService
}

func NewProductHandlers(productService ports.ProductService) *ProductHandlers {
	return &ProductHandlers{
		productService: productService,
	}
}

func (ph *ProductHandlers) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	response, err := ph.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandlers) GetRandomProductFromUserSsd(w http.ResponseWriter, r *http.Request) {
	response, err := ph.productService.GetRandomProductFromUserSsd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandlers) GetRandomProduct(w http.ResponseWriter, r *http.Request) {
	response, err := ph.productService.GetRandomProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandlers) CreateProductOrder(w http.ResponseWriter, r *http.Request) {
	var request ports.CreateProductOrderRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ph.productService.CreateProductOrder(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}
