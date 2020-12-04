package products

import (
	"context"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type delete struct {
	 repository rel.Repository
}

func (d delete) Delete(ctx context.Context, product *Product) error {
	if err := d.repository.Delete(ctx, product); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	return nil
}
