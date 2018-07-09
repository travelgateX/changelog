package resolver

import (
	"changelog/model"
	"changelog/scalars/version"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// ReleaseDataResolver :
type ReleaseDataResolver struct {
	Release *model.ReleaseData
	Error   *model.Error
}

// Code :
func (r *ReleaseDataResolver) Code() graphql.ID {
	return r.Release.Code
}

// Name :
func (r *ReleaseDataResolver) Name() *string {
	return r.Release.Name
}

// Version :
func (r *ReleaseDataResolver) Version() version.Version {
	return r.Release.Version
}

// ReleaseDate :
func (r *ReleaseDataResolver) ReleaseDate() time.Time {
	return r.Release.ReleaseDate
}

// CreatedAt :
func (r *ReleaseDataResolver) CreatedAt() time.Time {
	return r.Release.CreatedAt
}

// UpdatedAt :
func (r *ReleaseDataResolver) UpdatedAt() time.Time {
	return r.Release.UpdatedAt
}
