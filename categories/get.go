package categories

import (
	"context"

	"github.com/go-rel/rel"
)

type get struct {
	repository rel.Repository
}

func (g get) Get(ctx context.Context, category *[]Category) error {
	g.repository.MustFindAll(ctx, category)

	return nil
}
