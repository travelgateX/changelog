package resolver

import (
	"context"
)

// CreateCommitMutationArgs : create mutation
type CreateCommitMutationArgs struct {
	Message  string
	User     *string
	Release  *string
	Resource *string
	Category *string
}

// CreateCommit creates new commit in data base
func (r Resolver) CreateCommit(ctx context.Context, args CreateCommitMutationArgs) (*CommitResolver, error) {
	return &CommitResolver{}, nil
}
