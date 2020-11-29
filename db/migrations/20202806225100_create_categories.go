package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateCategories definition
func MigrateCreateCategories(schema *rel.Schema) {
	schema.CreateTable("categories", func(t *rel.Table) {
		t.ID("id")
		t.String("category_name")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		// t.Int("order")
	})

	// schema.CreateIndex("categories", "order", []string{"order"})
}

// RollbackCreateCategories definition
func RollbackCreateCategories(schema *rel.Schema) {
	schema.DropTable("categories")
}
