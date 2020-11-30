package products

import (
	"context"

	"github.com/go-rel/rel"
)

type get struct {
	repository rel.Repository
}

func (g get) Get(ctx context.Context, products *[]Product) error {
	g.repository.FindAll(ctx, products)

	return nil
}
