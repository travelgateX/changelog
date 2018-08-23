package resolver

import (
	"changelog/model"
	"fmt"
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
	fmt.Println("Intentando crear dato en BBDD")
	r.db.Create(&c)
	fmt.Println("Dato supuestamente creado")
	return &ChangeResolver{change: &c, err: nil}
}
