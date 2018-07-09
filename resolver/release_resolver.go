package resolver

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// ReleaseResolver :
type ReleaseResolver struct {
	Release *model.Release
	Error   *model.Error
}

// Code :
func (r *ReleaseResolver) Code() graphql.ID {
	return r.Release.ReleaseData.Code
}

// ReleaseData :
func (r *ReleaseResolver) ReleaseData() *ReleaseDataResolver {
	return &ReleaseDataResolver{r.Release.ReleaseData, r.Error}
}

// CreatedAt :
func (r *ReleaseResolver) CreatedAt() time.Time {
	return r.Release.ReleaseData.CreatedAt
}

// UpdatedAt :
func (r *ReleaseResolver) UpdatedAt() time.Time {
	return r.Release.ReleaseData.UpdatedAt
}
