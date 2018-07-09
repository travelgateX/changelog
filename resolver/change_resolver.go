package resolver

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// ChangeResolver :
type ChangeResolver struct {
	Change *model.ChangeData
	Error  *model.Error
}

// Code :
func (r *ChangeResolver) Code() graphql.ID {
	return r.Change.Code
}

// ChangeData :
func (r *ChangeResolver) ChangeData() *ChangeDataResolver {
	return &ChangeDataResolver{r.Change, r.Error}
}

// CreatedAt :
func (r *ChangeResolver) CreatedAt() time.Time {
	return r.Change.CreatedAt
}

// UpdatedAt :
func (r *ChangeResolver) UpdatedAt() time.Time {
	return r.Change.UpdatedAt
}
