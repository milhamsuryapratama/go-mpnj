package categories

import (
	"context"
	"strconv"
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
		category   = Category{ID: "1", CategoryName: "Makanan"}
		id, _      = strconv.Atoi(category.ID)
	)

	repository.ExpectFind(
		where.Eq("id", id),
	).Result(category)
	repository.ExpectDelete().For(&category)

	assert.NotPanics(t, func() {
		service.Delete(ctx, id)
	})

	repository.AssertExpectations(t)
}
