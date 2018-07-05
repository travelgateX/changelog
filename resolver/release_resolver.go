package resolver

import (
	"changelog/model"
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// ReleaseResolver :
type ReleaseResolver struct {
	Release *model.Release
	Error   *model.Error
}

// Code :
func (r *ReleaseResolver) Code() graphql.ID {
	return r.Release.Code
}

// ReleaseData :
func (r *ReleaseResolver) ReleaseData() *ReleaseDataResolver {
	return &ReleaseDataResolver{r.Release.ReleaseData, r.Error}
}

// CreatedAt :
func (r *ReleaseResolver) CreatedAt() datetime.DateTime {
	return r.Release.CreatedAt
}

// UpdatedAt :
func (r *ReleaseResolver) UpdatedAt() datetime.DateTime {
	return r.Release.UpdatedAt
}
