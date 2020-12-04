package products

import (
	"go-mpnj/categories"
	"time"
)

// Product ...
type Product struct {
	ID           int                  `json:"id"`
	ProductName  string               `json:"product_name"`
	Slug         string               `json:"slug"`
	Weight       int                  `json:"weight"`
	CapitalPrice int                  `json:"capital_price"`
	SellingPrice int                  `json:"selling_price"`
	Discount     int                  `json:"discount"`
	Stock        int                  `json:"stock"`
	Notes        string               `json:"notes"`
	Wishlist     int                  `json:"wishlist"`
	Sold         int                  `json:"sold"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	Category     categories.Category `json:"category" ref:"category_id" fk:"id"`
	CategoryID   int                  `json:"category_id"`
}
