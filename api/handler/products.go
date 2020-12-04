package handler

import (
	"encoding/json"
	"fmt"
	"go-mpnj/products"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Products ...
type Products struct {
	*chi.Mux
	products products.Service
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

// Create ...
func (p Products) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		product products.Product
	)

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := p.products.Create(ctx, &product); err != nil {
		render(w, err, 422)
		return
	}

	w.Header().Set("Location", fmt.Sprint(r.RequestURI, "/", product.ID))
	render(w, product, 201)
}

// NewProducts ...
func NewProducts(products products.Service) Products {
	handler := Products{
		Mux:      chi.NewMux(),
		products: products,
	}

	handler.Get("/", handler.Index)
	handler.Post("/", handler.Create)

	return handler
}
