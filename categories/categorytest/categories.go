package categorytest

import (
	context "context"
	categories "go-mpnj/categories"

	mock "github.com/stretchr/testify/mock"
)

// MockFunc ...
type MockFunc func(service *Service)

// Mock ...
func Mock(service *Service, funcs ...MockFunc) {
	for i := range funcs {
		if funcs[i] != nil {
			funcs[i](service)
		}
	}
}

// MockGet ...
func MockGet(result []categories.Category, err error) MockFunc {
	return func(service *Service) {
		service.On("Get", mock.Anything, mock.Anything).Return(func(ctx context.Context, output *[]categories.Category) error {
			*output = result
			return err
		})
	}
}

// MockCreate ...
func MockCreate(result categories.Category, err error) MockFunc {
	return func(service *Service) {
		service.On("Create", mock.Anything, mock.Anything).Return(func(ctx context.Context, output *categories.Category) error {
			*output = result
			return err
		})
	}
}

// MockDelete ...
func MockDelete() MockFunc {
	return func(service *Service) {
		service.On("Delete", mock.Anything, mock.Anything)
	}
}
