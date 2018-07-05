package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// SourceResolver :
type SourceResolver struct {
	Source *model.Source
	Error  *model.Error
}

// Code :
func (r *SourceResolver) Code() graphql.ID {
	return r.Source.Code
}

// Platform :
func (r *SourceResolver) Platform() *model.SourcePlatform {
	return &r.Source.Platform
}

// Author :
func (r *SourceResolver) Author() *string {
	return &r.Source.Author
}
