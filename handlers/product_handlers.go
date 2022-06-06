package handlers

import (
	"encoding/json"
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

func (ph *ProductHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := ph.productService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

type GetRandomProductFromUserSsdResponse struct {
	UserId  int    `json:"user_id"`
	SsdId   int    `json:"ssd_id"`
	Barcode string `json:"barcode"`
}

func (ph *ProductHandlers) GetRandomProductFromUserSsd(w http.ResponseWriter, r *http.Request) {
	userId, ssdId, barcode, err := ph.productService.GetRandomProductFromUserSsd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &GetRandomProductFromUserSsdResponse{
		UserId:  userId,
		SsdId:   ssdId,
		Barcode: barcode,
	}

	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandlers) Random(w http.ResponseWriter, r *http.Request) {
	product, err := ph.productService.Random()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
