package resolver

import (
	"context"

	"github.com/jinzhu/gorm"
)

// QueryResolver : is the entry point for all top-level read operations.
type QueryResolver struct {
	db *gorm.DB
}

// ChangeQueryArgs : are the arguments for the "Change" query.
type ChangeQueryArgs struct {
	First *int
}

// Changes : resolves a list of all changes
func (r QueryResolver) Changes(ctx context.Context, args ChangeQueryArgs) (*[]*ChangeResolver, error) {
	// ========= TODO: return something nice
	return nil, nil
}
