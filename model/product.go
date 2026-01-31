package model

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
	CategoryName string    `json:"category_name,omitempty"`
	Category     *Category `json:"category,omitempty"`
}
