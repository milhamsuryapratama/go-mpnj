package categories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, category *Category, id int) error {
	var categori Category
	if err := u.repository.Find(ctx, &categori, where.Eq("id", id)); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	categori.CategoryName = category.CategoryName
	if err := u.repository.Update(ctx, &categori); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	return nil
}
