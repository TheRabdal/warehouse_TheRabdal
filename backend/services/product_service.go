package services

import (
	"database/sql"
	"log"
	"warehouse/backend/models"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) GetProducts(statusFilter string) ([]models.Product, error) {
	var products []models.Product
	var rows *sql.Rows
	var err error

	if statusFilter != "" {
		rows, err = s.db.Query("SELECT sku, name, quantity, location, status FROM products WHERE status = ?", statusFilter)
	} else {
		rows, err = s.db.Query("SELECT sku, name, quantity, location, status FROM products")
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.SKU, &p.Name, &p.Quantity, &p.Location, &p.Status); err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	if err := ValidateProduct(p); err != nil {
		return err
	}

	stmt, err := s.db.Prepare("INSERT INTO products(sku, name, quantity, location, status) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(p.SKU, p.Name, p.Quantity, p.Location, p.Status)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(sku string, p *models.Product) error {
	if err := ValidateProduct(p); err != nil {
		return err
	}

	stmt, err := s.db.Prepare("UPDATE products SET name=?, quantity=?, location=?, status=? WHERE sku=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(p.Name, p.Quantity, p.Location, p.Status, sku)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(sku string) error {
	stmt, err := s.db.Prepare("DELETE FROM products WHERE sku=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(sku)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
