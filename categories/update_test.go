package categories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
)

func TestUpdate(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		category   = Category{ID: "1", CategoryName: "Makanan"}
		changes    = rel.NewChangeset(&category)
	)

	category.CategoryName = "Elektronik"

	repository.ExpectUpdate(changes).ForType("categories.Category")

	assert.NotPanics(t, func() {
		service.Update(ctx, &category, changes)
	})
	assert.Equal(t, category.CategoryName, "Elektronik")

	repository.AssertExpectations(t)
}
