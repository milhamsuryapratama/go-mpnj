package categories

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, category *Category, changes rel.Changeset) error {
	if err := category.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	u.repository.Update(ctx, category, changes)

	return nil
}
