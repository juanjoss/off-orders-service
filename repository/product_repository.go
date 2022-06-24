package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/juanjoss/off-orders-service/ports"
	_ "github.com/lib/pq"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository() *ProductRepository {
	host := os.Getenv("DB_HOST")
	driver := os.Getenv("DB_DRIVER")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, dbPort, dbUser, dbPassword, dbName, sslMode)

	db, err := sqlx.Connect(driver, source)
	if err != nil {
		log.Fatalf("unable to connect to DB: %v", err.Error())
	}

	repo := &ProductRepository{
		db: db,
	}

	return repo
}

/*
	Returns all products.
*/
func (pr *ProductRepository) GetAllProducts() (ports.GetAllProductsResponse, error) {
	response := ports.GetAllProductsResponse{}

	err := pr.db.Select(&response.Products, `SELECT * FROM products`)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
	Returns a random product from a user's ssd.
*/
func (pr *ProductRepository) GetRandomProductFromUserSsd() (ports.GetRandomProductFromUserSsdResponse, error) {
	response := ports.GetRandomProductFromUserSsdResponse{}

	rows, err := pr.db.Query(`
		SELECT ssds.id AS ssdId, product_ssds.barcode AS barcode
		FROM (users JOIN ssds ON users.id = ssds.id) JOIN product_ssds ON ssds.id = product_ssds.ssd_id
		ORDER BY RANDOM() 
		LIMIT 1
	`)
	if err != nil {
		return response, err
	}

	for rows.Next() {
		err = rows.Scan(&response.SsdId, &response.Barcode)
		if err != nil {
			return response, err
		}
	}
	if err = rows.Err(); err != nil {
		return response, err
	}

	return response, nil
}

/*
	Returns a random product.
*/
func (pr *ProductRepository) GetRandomProduct() (ports.GetRandomProductResponse, error) {
	response := ports.GetRandomProductResponse{}

	err := pr.db.Get(&response, `SELECT * FROM products ORDER BY RANDOM() LIMIT 1`)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
	Creates a product order.
*/
func (pr *ProductRepository) CreateProductOrder(request ports.CreateProductOrderRequest) error {
	_, err := pr.db.Exec(
		`INSERT INTO product_orders (ssd_id, barcode, timestamp, quantity, status) 
		VALUES ($1, $2, NOW(), $3, 'pending')
		ON CONFLICT DO NOTHING`,
		request.SsdId, request.Barcode, request.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}
