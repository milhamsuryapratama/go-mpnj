package handler

import (
	"go-mpnj/products"
	"net/http"

	"github.com/go-chi/chi"
)

// Products ...
type Products struct {
	*chi.Mux
	products products.Service
}

type yolo struct {
	Nama   string
	Alamat string
}

// Index ...
func (p Products) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		result []products.Product
	)

	p.products.Get(ctx, &result)
	render(w, result, 200)
}

// NewProducts ...
func NewProducts(products products.Service) Products {
	handler := Products{
		Mux:      chi.NewMux(),
		products: products,
	}

	handler.Get("/", handler.Index)

	return handler
}
