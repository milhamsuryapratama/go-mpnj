package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-rel/rel"
	"go-mpnj/products"
	"net/http"
	"strconv"

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

func (p Products) Destroy(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		product = ctx.Value(loadKey).(products.Product)
	)

	if err := p.products.Delete(ctx, &product); err != nil {
		render(w, ErrBadRequest, 400)
		return
	}

	render(w, nil, 204)
}

func (p Products) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx 	= r.Context()
		product = ctx.Value(loadKey).(products.Product)
		changes = rel.NewChangeset(&product)
	)

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := p.products.Update(ctx, &product, changes); err != nil {
		render(w, err, 422)
		return
	}

	render(w, product, 200)
}

func (p Products) Load(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			id, _ = strconv.Atoi(chi.URLParam(r, "ID"))
			product products.Product
		)

		err := p.products.FindByID(ctx, &product, uint(id))
		if err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				render(w, err, 404)
				return
			}

			panic(err)
		}

		ctx = context.WithValue(ctx, loadKey, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewProducts ...
func NewProducts(products products.Service) Products {
	handler := Products{
		Mux:      chi.NewMux(),
		products: products,
	}

	handler.Get("/", handler.Index)
	handler.Post("/", handler.Create)
	handler.With(handler.Load).Delete("/{ID}", handler.Destroy)
	handler.Put("/{ID}", handler.Update)

	return handler
}
