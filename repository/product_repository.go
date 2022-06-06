package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/juanjoss/off-orders-service/model"
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
func (pr *ProductRepository) GetAll() ([]*model.Product, error) {
	products := []*model.Product{}

	err := pr.db.Select(&products, `SELECT * FROM products`)
	if err != nil {
		return products, err
	}

	return products, nil
}

/*
	Returns a random product from a user's ssd.
*/
func (pr *ProductRepository) GetRandomProductFromUserSsd() (int, int, string, error) {
	var userId, ssdId int
	var barcode string

	rows, err := pr.db.Query(`
		SELECT users.id AS userId, ssds.id AS ssdId, product_ssds.barcode AS barcode
		FROM (users JOIN ssds ON users.id = ssds.id) JOIN product_ssds ON ssds.id = product_ssds.ssd_id
		ORDER BY RANDOM() 
		LIMIT 1
	`)
	if err != nil {
		return userId, ssdId, barcode, err
	}

	for rows.Next() {
		err = rows.Scan(&userId, &ssdId, &barcode)
		if err != nil {
			return userId, ssdId, barcode, err
		}
	}
	if err = rows.Err(); err != nil {
		return userId, ssdId, barcode, err
	}

	return userId, ssdId, barcode, nil
}

/*
	Returns a random product.
*/
func (pr *ProductRepository) Random() (*model.Product, error) {
	product := &model.Product{}

	rows, err := pr.db.Queryx(
		`SELECT * FROM products ORDER BY RANDOM() LIMIT 1`,
	)
	if err != nil {
		return product, err
	}

	for rows.Next() {
		err = rows.StructScan(product)
		if err != nil {
			return product, err
		}
	}
	if err = rows.Err(); err != nil {
		return product, err
	}

	return product, nil
}
