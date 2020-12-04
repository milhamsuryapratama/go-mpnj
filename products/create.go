package products

import (
	"context"

	"github.com/go-rel/rel"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, product *Product) error {
	if err := c.repository.Insert(ctx, product); err != nil {
		return err
	}

	return nil
}
