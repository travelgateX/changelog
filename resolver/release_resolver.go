package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// ReleaseResolver :
type ReleaseResolver struct {
	release *model.Release
	err     *model.AdviseMessageData
}

// Code :
func (r *ReleaseResolver) Code() graphql.ID {
	return r.release.ReleaseData.Code
}

// ReleaseData :
func (r *ReleaseResolver) ReleaseData() *ReleaseDataResolver {
	return &ReleaseDataResolver{r.release.ReleaseData, r.err}
}

// CreatedAt :
func (r *ReleaseResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.release.ReleaseData.CreatedAt}
}

// UpdatedAt :
func (r *ReleaseResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.release.ReleaseData.UpdatedAt}
}
