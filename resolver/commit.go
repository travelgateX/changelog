package resolver

import (
	"changelog/model"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

// CommitResolver : resolves Change type
type CommitResolver struct {
	commit model.Commit
}

// ID unique Commit identifier -> fmt.Sprintf (int32 -> string)
func (r *CommitResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.commit.ID))
}

// Message : resolves commit message
func (r *CommitResolver) Message() string {
	return r.commit.Message
}

// User : resolves commit user
func (r *CommitResolver) User() *string {
	return &r.commit.User
}

// Release : resolves commit release
func (r *CommitResolver) Release() *string {
	return &r.commit.Release
}

// Resource : resolves commit release
func (r *CommitResolver) Resource() *string {
	return &r.commit.Resource
}

// Category : resolves commit category
func (r *CommitResolver) Category() model.Category {
	return r.commit.Category
}

// CreatedAt : resolves commit created_at
func (r *CommitResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.commit.CreatedAt}
}
