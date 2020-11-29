package categories

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, category *Category) error {
	if err := category.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	c.repository.Insert(ctx, category)
	return nil
}
