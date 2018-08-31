package resolver

import (
	"changelog/model"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// ChangeDataResolver :
type ChangeDataResolver struct {
	db         *gorm.DB
	changeData *model.ChangeData
}

// Code :
func (r *ChangeDataResolver) Code() graphql.ID {
	return r.changeData.Code
}

// Message :
func (r *ChangeDataResolver) Message() string {
	return r.changeData.Message
}

// Project :
func (r *ChangeDataResolver) Project() string {
	fmt.Println("ChangeDataResolver : Project")
	return *r.changeData.Project
}

// ChangeDate :
func (r *ChangeDataResolver) ChangeDate() graphql.Time {
	return graphql.Time{Time: r.changeData.ChangeDate}
}

// Type :
func (r *ChangeDataResolver) Type() model.ChangeType {
	return r.changeData.Type
}

// Release :
func (r *ChangeDataResolver) Release() *ReleaseResolver {
	var releaseData model.ReleaseData
	r.db.First(&releaseData, r.changeData.ReleaseID)
	return &ReleaseResolver{release: releaseData.ReleaseObject(), err: nil}
}

// Source :
func (r *ChangeDataResolver) Source() *SourceResolver {
	var source model.Source
	r.db.First(&source, r.changeData.SourceID)
	return &SourceResolver{source: &source, err: nil}
}

// CreatedAt :
func (r *ChangeDataResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.changeData.CreatedAt}
}

// UpdatedAt :
func (r *ChangeDataResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.changeData.UpdatedAt}
}
