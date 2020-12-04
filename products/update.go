package products

import (
	"context"
	"github.com/go-rel/rel"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, product *Product, changes rel.Changeset) error {
	u.repository.Update(ctx, product, changes)
	return nil
}
