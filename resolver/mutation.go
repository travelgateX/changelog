package resolver

import (
	"changelog/model"
	"context"

	"github.com/jinzhu/gorm"
)

// MutationResolver :
type MutationResolver struct {
	db *gorm.DB
}

// CreateChangeMutationArgs : are the arguments for the "CreateChange" mutation.
type CreateChangeMutationArgs struct {
	Body *string
}

// CreateChange : resolves a list of all changes
func (r MutationResolver) CreateChange(ctx context.Context, args CreateChangeMutationArgs) (*model.Change, error) {
	// ========= TODO: return something nice
	return nil, nil
}
