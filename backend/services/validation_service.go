package services

import (
	"errors"
	"warehouse/backend/models"
)

func ValidateProduct(p *models.Product) error {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if p.SKU == "" {
		return errors.New("product SKU cannot be empty")
	}
	if p.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	return nil
}
