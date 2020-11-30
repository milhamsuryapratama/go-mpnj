package migrations

import "github.com/go-rel/rel"

// MigrateCreateProducts ...
func MigrateCreateProducts(schema *rel.Schema) {
	schema.CreateTable("products", func(t *rel.Table) {
		t.ID("id")
		t.String("product_name")
		t.String("slug")
		t.String("weight")
		t.Int("capital_price")
		t.Int("selling_price")
		t.Int("discount")
		t.Int("stock")
		t.Text("notes")
		t.Int("wishlist")
		t.Int("sold")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.Int("category_id", rel.Unsigned(true))
		t.ForeignKey("category_id", "categories", "id")
	})
}

// RollbackCreateProducts ...
func RollbackCreateProducts(schema *rel.Schema) {
	schema.DropTable("products")
}
