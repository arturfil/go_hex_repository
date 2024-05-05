package types

import "time"

type ProductStore interface {
	GetProducts() ([]*Product, error)
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Descrtipion string    `json:"descrtipion"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"` // TODO: Better Structure to this
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
