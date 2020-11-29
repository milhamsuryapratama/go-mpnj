package categories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, id int) error {
	var category Category
	if err := d.repository.Find(ctx, &category, where.Eq("id", id)); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	if err := d.repository.Delete(ctx, &category); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	return nil
}
