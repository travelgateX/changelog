package resolver

import (
	"changelog/model"
	"changelog/relay"
	"context"
)

// Changelog resolves a list of changes
func (r *Resolver) Changelog(ctx context.Context, args relay.Filters) (*ChangeConnectionResolver, error) {
	var changes []*model.ChangeData

	r.db.Find(&changes)
	changesInterface := changesToSliceInterface(changes)
	filterArgs := relay.NewConnectionArguments(args.Filters())
	conn := relay.ConnectionFromArray(changesInterface, filterArgs)

	return &ChangeConnectionResolver{db: r.db, conn: conn}, nil
}

func changesToSliceInterface(changes []*model.ChangeData) []interface{} {
	iFace := []interface{}{}
	for _, todo := range changes {
		iFace = append(iFace, todo)
	}
	return iFace
}
