package products

import (
	"context"

	"github.com/go-rel/rel"
)

// Service ...
type Service interface {
	Get(ctx context.Context, products *[]Product) error
	Create(ctx context.Context, product *Product) error
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
