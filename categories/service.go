package categories

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "categories")))
)

// Service ...
type Service interface {
	Get(ctx context.Context, category *[]Category) error
	Create(ctx context.Context, category *Category) error
}

type service struct {
	get
	create
}

var _ Service = (*service)(nil)

// New ...
func New(repository rel.Repository) Service {
	return service{
		get:    get{repository: repository},
		create: create{repository: repository},
	}
}
