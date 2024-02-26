package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Producer interface {
	List() ([]Product, error)
	Get(int) (Product, error)
	Create(Product) (Product, error)
	Update(Product) (Product, error)
	Delete(Product) error
}
