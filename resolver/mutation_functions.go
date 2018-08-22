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
	r.db.Save(&c)
	return &ChangeResolver{
		change: &c,
		err:    nil,
	}
}
