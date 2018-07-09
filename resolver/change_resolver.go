package resolver

import (
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// ChangeResolver :
type ChangeResolver struct {
	db     *gorm.DB
	change *model.ChangeData
	err    *[]*model.Error
}

// Code :
func (r *ChangeResolver) Code() graphql.ID {
	return r.change.Code
}

// ChangeData :
func (r *ChangeResolver) ChangeData() *ChangeDataResolver {
	return &ChangeDataResolver{db: r.db, changeData: r.change, err: r.err}
}

// CreatedAt :
func (r *ChangeResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.change.CreatedAt}
}

// UpdatedAt :
func (r *ChangeResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.change.UpdatedAt}
}
