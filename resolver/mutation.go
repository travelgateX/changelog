package resolver

import (
	"changelog/model"
	"context"
	"time"
)

// CreateCommitMutationArgs : create mutation
type CreateCommitMutationArgs struct {
	Message  string
	User     *string
	Release  *string
	Resource *string
	Category model.Category
}

// CreateCommit creates new commit in data base
func (r Resolver) CreateCommit(ctx context.Context, args CreateCommitMutationArgs) (*CommitResolver, error) {
	commit := model.Commit{
		Message:   args.Message,
		User:      getOptionalStrValue(args.User),
		Release:   getOptionalStrValue(args.Release),
		Resource:  getOptionalStrValue(args.Resource),
		Category:  args.Category,
		CreatedAt: time.Now(),
	}
	r.db.Create(&commit)

	return &CommitResolver{commit: commit}, nil
}
