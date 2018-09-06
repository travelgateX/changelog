package resolver

import (
	"changelog/model"
	"changelog/scalars/version"

	graphql "github.com/graph-gophers/graphql-go"
)

// ReleaseDataResolver :
type ReleaseDataResolver struct {
	Release *model.ReleaseData
	Error   *model.AdviseMessageData
}

// Code :
func (r *ReleaseDataResolver) Code() graphql.ID {
	return r.Release.Code
}

// Name :
func (r *ReleaseDataResolver) Name() *string {
	return &r.Release.Name
}

// Version :
func (r *ReleaseDataResolver) Version() version.Version {
	return r.Release.Version
}

// ReleaseDate :
func (r *ReleaseDataResolver) ReleaseDate() graphql.Time {
	return graphql.Time{Time: r.Release.ReleaseDate}
}

// CreatedAt :
func (r *ReleaseDataResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.Release.CreatedAt}
}

// UpdatedAt :
func (r *ReleaseDataResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.Release.UpdatedAt}
}
