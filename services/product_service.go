package services

import (
	"github.com/juanjoss/off-orders-service/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (ps *ProductService) GetAllProducts() (ports.GetAllProductsResponse, error) {
	response, err := ps.repo.GetAllProducts()
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ps *ProductService) GetRandomProductFromUserSsd() (ports.GetRandomProductFromUserSsdResponse, error) {
	response, err := ps.repo.GetRandomProductFromUserSsd()
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ph *ProductService) GetRandomProduct() (ports.GetRandomProductResponse, error) {
	response, err := ph.repo.GetRandomProduct()
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ph *ProductService) CreateProductOrder(request ports.CreateProductOrderRequest) error {
	err := ph.repo.CreateProductOrder(request)
	if err != nil {
		return err
	}

	return nil
}
