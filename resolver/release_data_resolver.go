package resolver

import (
	"changelog/model"
	"changelog/scalars/datetime"
	"changelog/scalars/version"

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

// Version :
func (r *ReleaseDataResolver) Version() version.Version {
	return r.Release.Version
}

// ReleaseDate :
func (r *ReleaseDataResolver) ReleaseDate() datetime.DateTime {
	return r.Release.ReleaseDate
}
