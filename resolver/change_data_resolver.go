package resolver

import (
	"changelog/model"
	"changelog/scalars/datetime"

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
func (r *ChangeDataResolver) Project() *string {
	return r.ChangeData.Project
}

// ChangeDate :
func (r *ChangeDataResolver) ChangeDate() datetime.DateTime {
	return r.ChangeData.ChangeDate
}

// Type :
func (r *ChangeDataResolver) Type() model.ChangeType {
	return r.ChangeData.Type
}

// Release :
func (r *ChangeDataResolver) Release() model.Release {
	return r.ChangeData.Release
}

// Source :
func (r *ChangeDataResolver) Source() model.Source {
	return r.ChangeData.Source
}

// CreatedAt :
func (r *ChangeDataResolver) CreatedAt() datetime.DateTime {
	return r.ChangeData.CreatedAt
}

// UpdatedAt :
func (r *ChangeDataResolver) UpdatedAt() datetime.DateTime {
	return r.ChangeData.UpdatedAt
}
