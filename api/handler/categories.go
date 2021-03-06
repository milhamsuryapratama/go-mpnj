package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-mpnj/categories"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type ctx int

const (
	bodyKey ctx = 0
	loadKey ctx = 1
)

// Categories ...
type Categories struct {
	*chi.Mux
	categories categories.Service
}

// Index ...
func (c Categories) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		result []categories.Category
	)

	c.categories.Get(ctx, &result)
	render(w, result, 200)
}

// Create ...
func (c Categories) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		category categories.Category
	)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := c.categories.Create(ctx, &category); err != nil {
		render(w, err, 422)
		return
	}

	w.Header().Set("Location", fmt.Sprint(r.RequestURI, "/", category.ID))
	render(w, category, 201)
}

// Update ...
func (c Categories) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		category = ctx.Value(loadKey).(categories.Category)
		changes  = rel.NewChangeset(&category)
	)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := c.categories.Update(ctx, &category, changes); err != nil {
		render(w, err, 422)
		return
	}

	render(w, category, 200)
}

// Destroy ...
func (c Categories) Destroy(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		category = ctx.Value(loadKey).(categories.Category)
	)

	if err := c.categories.Delete(ctx, &category); err != nil {
		render(w, ErrBadRequest, 400)
		return
	}

	render(w, nil, 204)
}

func (c Categories) Load(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx 	= r.Context()
			id, _ 	= strconv.Atoi(chi.URLParam(r, "ID"))
			category categories.Category
		)

		err := c.categories.FindByID(ctx, &category, uint(id))
		if err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				render(w, err, 404)
				return
			}
			panic(err)
		}

		ctx = context.WithValue(ctx, loadKey, category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewCategories ...
func NewCategories(categories categories.Service) Categories {
	handler := Categories{
		Mux:        chi.NewMux(),
		categories: categories,
	}

	handler.Get("/", handler.Index)
	handler.Post("/", handler.Create)
	handler.With(handler.Load).Put("/{ID}", handler.Update)
	handler.With(handler.Load).Delete("/{ID}", handler.Destroy)

	return handler
}
