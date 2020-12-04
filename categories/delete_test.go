package categories

import (
	"context"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		category   = Category{ID: 1, CategoryName: "Makanan"}
	)

	repository.ExpectFind(
		where.Eq("id", category.ID),
	).Result(category)
	repository.ExpectDelete().For(&category)

	assert.NotPanics(t, func() {
		service.Delete(ctx, category.ID)
	})

	repository.AssertExpectations(t)
}
