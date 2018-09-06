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
	err    *[]model.AdviseMessageData
}

// Code :
func (r *ChangeResolver) Code() graphql.ID {
	if r.err != nil {
		return ""
	}
	return r.change.Code
}

// ChangeData :
func (r *ChangeResolver) ChangeData() *ChangeDataResolver {
	if r.change == nil {
		return nil
	}
	return &ChangeDataResolver{db: r.db, changeData: r.change}
}

// AdviseMessage :
func (r *ChangeResolver) AdviseMessage(args struct{ Level *[]*model.AdviseMessageLevel }) *[]*AdviseMessageResolver {
	if r.err == nil {
		return nil
	}
	s := []*AdviseMessageResolver{}
	for _, res := range *r.err {
		s = append(s, &AdviseMessageResolver{AdviseMessage: &res})
	}

	return &s
}

// CreatedAt :
func (r *ChangeResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.change.CreatedAt}
}

// UpdatedAt :
func (r *ChangeResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.change.UpdatedAt}
}
