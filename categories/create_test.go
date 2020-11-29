package categories

import (
	"context"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		category   = Category{CategoryName: "Elektronik"}
	)

	repository.ExpectInsert().For(&category)

	assert.Nil(t, service.Create(ctx, &category))

	repository.AssertExpectations(t)
}
