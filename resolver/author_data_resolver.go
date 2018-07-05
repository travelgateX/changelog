package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// AuthorDataResolver :
type AuthorDataResolver struct {
	Author *model.AuthorData
	Error  *model.Error
}

// Code :
func (r *AuthorDataResolver) Code() graphql.ID {
	return r.Author.Code
}

// Name :
func (r *AuthorDataResolver) Name() *string {
	return &r.Author.Name
}
