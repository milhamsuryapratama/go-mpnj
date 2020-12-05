package categories

import (
	"context"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		category   = Category{ID: 1, CategoryName: "Makanan"}
	)

	repository.ExpectDelete().ForType("categories.Category")

	assert.NotPanics(t, func() {
		service.Delete(ctx, &category)
	})

	repository.AssertExpectations(t)
}
