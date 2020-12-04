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
	Update(ctx context.Context, category *Category, changes rel.Changeset) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, category *Category, id uint) error
}

type service struct {
	get
	create
	update
	delete
	findByID
}

var _ Service = (*service)(nil)

// New ...
func New(repository rel.Repository) Service {
	return service{
		get:    get{repository: repository},
		create: create{repository: repository},
		update: update{repository: repository},
		delete: delete{repository: repository},
		findByID: findByID{repository: repository},
	}
}
