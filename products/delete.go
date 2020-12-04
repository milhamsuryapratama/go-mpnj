package products

import (
	"context"
	"fmt"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

type delete struct {
	 repository rel.Repository
}

func (d delete) Delete(ctx context.Context, id int) error {
	var product Product
	if err := d.repository.Find(ctx, &product, where.Eq("id", id)); err != nil {
		fmt.Println("YOlo")
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	if err := d.repository.Delete(ctx, &product); err != nil {
		logger.Warn("Query error", zap.Error(err))
		return err
	}

	return nil
}
