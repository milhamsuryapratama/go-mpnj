package products

import (
	"context"
	"go.uber.org/zap"

	"github.com/go-rel/rel"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "products")))
)

// Service ...
type Service interface {
	Get(ctx context.Context, products *[]Product) error
	Create(ctx context.Context, product *Product) error
	Delete(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product, changes rel.Changeset) error
	FindByID(ctx context.Context, product *Product ,id uint) error
}

type service struct {
	get
	create
	delete
	update
	findByID
}

var _ Service = (*service)(nil)

// New ...
func New(repository rel.Repository) Service {
	return service{
		get:    get{repository: repository},
		create: create{repository: repository},
		delete: delete{repository: repository},
		update: update{repository: repository},
		findByID: findByID{repository: repository},
	}
}
