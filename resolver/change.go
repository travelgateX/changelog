package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// ChangeResolver : resolves Change type
type ChangeResolver struct {
	change model.Change
}

// ID Unique AccessConfiguration identifier
func (r *ChangeResolver) ID() graphql.ID {
	return graphql.ID(r.change.ID)
}

// Body resolves the change's unique identifier.
func (r *ChangeResolver) Body() *string {
	return &r.change.Body
}
