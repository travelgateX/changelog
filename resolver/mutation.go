package resolver

import (
	"context"

	"github.com/jinzhu/gorm"
)

// MutationResolver :
type MutationResolver struct {
	db *gorm.DB
}

// CreateCommitMutationArgs : commits request
type CreateCommitMutationArgs struct {
	Message  string
	User     *string
	Release  string
	Resource string
	Category string
}

// CreateCommit creates new commit in data base
func (r MutationResolver) CreateCommit(ctx context.Context, args CreateCommitMutationArgs) (*CommitResolver, error) {
	return &CommitResolver{}, nil
}
