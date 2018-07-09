package resolver

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// ChangeDataResolver :
type ChangeDataResolver struct {
	ChangeData *model.ChangeData
	Error      *model.Error
}

// Code :
func (r *ChangeDataResolver) Code() graphql.ID {
	return r.ChangeData.Code
}

// Message :
func (r *ChangeDataResolver) Message() string {
	return r.ChangeData.Message
}

// Project :
func (r *ChangeDataResolver) Project() string {
	return *r.ChangeData.Project
}

// ChangeDate :
func (r *ChangeDataResolver) ChangeDate() time.Time {
	return r.ChangeData.ChangeDate
}

// Type :
func (r *ChangeDataResolver) Type() model.ChangeType {
	return r.ChangeData.Type
}

// Release :
func (r *ChangeDataResolver) Release() *ReleaseResolver {
	return &ReleaseResolver{&r.ChangeData.Release, nil}
}

// Source :
func (r *ChangeDataResolver) Source() *SourceResolver {
	return &SourceResolver{&r.ChangeData.Source, nil}
}

// CreatedAt :
func (r *ChangeDataResolver) CreatedAt() time.Time {
	return r.ChangeData.CreatedAt
}

// UpdatedAt :
func (r *ChangeDataResolver) UpdatedAt() time.Time {
	return r.ChangeData.UpdatedAt
}
