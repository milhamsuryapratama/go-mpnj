package products

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type findByID struct {
	repository rel.Repository
}

func (f findByID) FindByID(ctx context.Context, product *Product ,id uint) error {
	if err := f.repository.Find(ctx, product, where.Eq("id", id)); err != nil {
		return err
	}
	return nil
}
