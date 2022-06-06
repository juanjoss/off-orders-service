package services

import (
	"github.com/juanjoss/off-orders-service/model"
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

func (ps *ProductService) GetAll() ([]*model.Product, error) {
	products, err := ps.repo.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (ps *ProductService) GetRandomProductFromUserSsd() (int, int, string, error) {
	userId, ssdId, barcode, err := ps.repo.GetRandomProductFromUserSsd()
	if err != nil {
		return userId, ssdId, barcode, err
	}

	return userId, ssdId, barcode, nil
}

func (ph *ProductService) Random() (*model.Product, error) {
	product, err := ph.repo.Random()
	if err != nil {
		return product, err
	}

	return product, nil
}
