package categories

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type findByID struct {
	repository rel.Repository
}

func (f findByID) FindByID(ctx context.Context, category *Category, id uint) error {
	if err := f.repository.Find(ctx, category, where.Eq("id", id)); err != nil {
		return err
	}
	return nil
}
