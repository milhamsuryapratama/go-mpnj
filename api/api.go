package api

import (
	"go-mpnj/api/handler"
	"go-mpnj/categories"

	"github.com/go-chi/chi"
	"github.com/go-rel/rel"
)

// NewMux ...
func NewMux(repository rel.Repository) *chi.Mux {
	var (
		mux               = chi.NewMux()
		categories        = categories.New(repository)
		categoriesHandler = handler.NewCategories(repository, categories)
	)

	mux.Mount("/categories", categoriesHandler)

	return mux
}
