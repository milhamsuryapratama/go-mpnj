package api

import (
	"go-mpnj/api/handler"
	"go-mpnj/categories"
	"go-mpnj/products"

	"github.com/go-chi/chi"
	"github.com/go-rel/rel"
)

// NewMux ...
func NewMux(repository rel.Repository) *chi.Mux {
	var (
		mux               = chi.NewMux()
		categories        = categories.New(repository)
		products          = products.New(repository)
		categoriesHandler = handler.NewCategories(repository, categories)
		productsHandler   = handler.NewProducts(products)
	)

	mux.Mount("/categories", categoriesHandler)
	mux.Mount("/products", productsHandler)

	return mux
}
