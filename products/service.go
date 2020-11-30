package products

import (
	"context"

	"github.com/go-rel/rel"
)

// Service ...
type Service interface {
	Get(ctx context.Context, products *[]Product) error
}

type service struct {
	get
}

var _ Service = (*service)(nil)

// New ...
func New(repository rel.Repository) Service {
	return service{
		get: get{repository: repository},
	}
}
