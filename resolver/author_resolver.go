package resolver

import (
	"changelog/model"
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// AuthorResolver :
type AuthorResolver struct {
	Author *model.Author
	Error  *model.Error
}

// Code :
func (r *AuthorResolver) Code() graphql.ID {
	return r.Author.Code
}

// AuthorData :
func (r *AuthorResolver) AuthorData() *AuthorDataResolver {
	return &AuthorDataResolver{&r.Author.AuthorData, nil}
}

// CreatedAt :
func (r *AuthorResolver) CreatedAt() datetime.DateTime {
	return r.Author.CreatedAt
}

// UpdatedAt :
func (r *AuthorResolver) UpdatedAt() datetime.DateTime {
	return r.Author.UpdatedAt
}
