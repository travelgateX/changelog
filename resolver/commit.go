package resolver

import (
	"changelog/model"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

// CommitResolver : resolves Change type
type CommitResolver struct {
	Commit model.Commit
}

// ID unique Commit identifier -> fmt.Sprintf (int32 -> string)
func (r *CommitResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.Commit.ID))
}

// Message : resolves commit message
func (r *CommitResolver) Message() string {
	return r.Commit.Message
}

// User : resolves commit user
func (r *CommitResolver) User() *string {
	return &r.Commit.User
}

// Release : resolves commit release
func (r *CommitResolver) Release() *string {
	return &r.Commit.Release
}

// Resource : resolves commit release
func (r *CommitResolver) Resource() *string {
	return &r.Commit.Resource
}

// Category : resolves commit category
func (r *CommitResolver) Category() model.Category {
	return r.Commit.Category
}

// Released : resolves commit released
func (r *CommitResolver) Released() bool {
	return r.Commit.Released
}

// CreatedAt : resolves commit created_at
func (r *CommitResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.Commit.CreatedAt}
}
