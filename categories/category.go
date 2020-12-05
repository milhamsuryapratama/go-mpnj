package categories

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	// CategoryPrefixURL ...
	CategoryPrefixURL = os.Getenv("URL") + "category/"
	//ErrorCategoryNameBlank ...
	ErrorCategoryNameBlank = errors.New("Category Name can't blank")
)

// Category ...
type Category struct {
	ID           int       			`json:"id"`
	CategoryName string    			`json:"category_name"`
	CreatedAt    time.Time 			`json:"created_at"`
	UpdatedAt    time.Time 			`json:"updated_at"`
	//Products     []products.Product `json:"products"`
}

// Validate ...
func (c Category) Validate() error {
	var err error
	switch {
	case len(c.CategoryName) == 0:
		err = ErrorCategoryNameBlank
	}

	return err
}

// MarshalJSON ...
func (c Category) MarshalJSON() ([]byte, error) {
	type Alias Category

	return json.Marshal(struct {
		Alias
		URL string `json:"url"`
	}{
		Alias: Alias(c),
		URL:   fmt.Sprint(CategoryPrefixURL, c.ID),
	})
}
