package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// SourceResolver :
type SourceResolver struct {
	source *model.Source
	err    *model.AdviseMessageData
}

// Code :
func (r *SourceResolver) Code() graphql.ID {
	return r.source.Code
}

// Platform :
func (r *SourceResolver) Platform() model.SourcePlatform {
	return r.source.Platform
}

// Author :
func (r *SourceResolver) Author() *string {
	return &r.source.Author
}
