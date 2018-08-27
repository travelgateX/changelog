package resolver

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// CreateChange :
func (r *Resolver) CreateChange(args struct{ Input *model.CreateChangeInput }) *ChangeResolver {

	t := time.Now()
	c := model.ChangeData{
		Code:       graphql.ID(*args.Input.Code),
		Message:    args.Input.Message,
		Project:    &args.Input.Project,
		ChangeDate: t,
		Type:       args.Input.Type,
		ReleaseID:  1,
		SourceID:   2,
	}
	r.db.Create(&c)
	return &ChangeResolver{change: &c, err: nil}
}

// UpdateChange :
func (r *Resolver) UpdateChange(args struct{ Input *model.UpdateChangeInput }) *ChangeResolver {

	t := time.Now()
	c := model.ChangeData{
		Code:       graphql.ID(args.Input.Code),
		Message:    args.Input.Message,
		Project:    &args.Input.Project,
		ChangeDate: t,
		Type:       args.Input.Type,
	}

	r.db.Model(&c).Updates(model.ChangeData{
		Message:    args.Input.Message,
		Project:    &args.Input.Project,
		ChangeDate: t,
		Type:       args.Input.Type})

	return &ChangeResolver{change: &c, err: nil}
}

// DeleteChange :
func (r *Resolver) DeleteChange(args struct{ Input *model.DeleteChangeInput }) *ChangeResolver {

	c := model.ChangeData{
		Code: graphql.ID(args.Input.Code),
	}

	// Get all matched records
	r.db.Where("code = ?", args.Input.Code).Delete(c)
	return &ChangeResolver{change: &c, err: nil}
}
