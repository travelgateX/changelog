package resolver

import (
	"changelog/model"
	"context"

	"github.com/jinzhu/gorm"
)

// QueryResolver :
type QueryResolver struct {
	db *gorm.DB
}

// CommitsQueryArgs : commits request
type CommitsQueryArgs struct {
	Resource *string
}

// Commits resolves a list of commits. If no arguments are provided, all commits are fetched.
func (r QueryResolver) Commits(ctx context.Context, args CommitsQueryArgs) (*[]*CommitResolver, error) {
	var commits []model.Commit

	// TODO: resolver going directly to data base... new layer here?
	r.db.Find(&commits)

	// TODO: load resolvers, handle errors
	var resolvers = make([]*CommitResolver, 0, len(commits))
	for _, commit := range commits {
		resolver := &CommitResolver{commit: commit}
		resolvers = append(resolvers, resolver)
	}

	return &resolvers, nil
}
