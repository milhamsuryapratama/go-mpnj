package categories

import (
	"context"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		categories []Category
		result     = []Category{{ID: 1, CategoryName: "Makanan"}}
	)

	repository.ExpectFindAll().Result(result)

	assert.NotPanics(t, func() {
		service.Get(ctx, &categories)
		assert.Equal(t, result, categories)
		assert.Equal(t, result[0].CategoryName, "Makanan")
	})

	repository.AssertExpectations(t)
}
