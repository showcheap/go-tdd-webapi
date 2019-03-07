package models

import (
	"errors"
)

// Product ...
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func (p *Product) getProduct() error {
	return errors.New("Not implemented")
}

func (p *Product) updateProduct() error {
	return errors.New("Not implemented")
}

func (p *Product) deleteProduct() error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct() error {
	return errors.New("Not implemented")
}

// GetProducts get all products
func GetProducts() ([]Product, error) {
	products = make([]Product, 0)
	return products, nil
}
