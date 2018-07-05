package resolver

import (
	"changelog/model"
	"context"
)

// ChangelogInput : commits request
type ChangelogInput struct {
	Filter *ChangelogFilterInput
}

// ChangelogFilterInput :
type ChangelogFilterInput struct {
	Type *model.ChangeType
}

// Changelog resolves a list of changes
func (r *Resolver) Changelog(ctx context.Context, args *ChangelogInput) (*ChangeConnectionResolver, error) {
	//var commits []model.Commit

	// TODO: resolver going directly to data base... new layer here?
	//r.db.Find(&commits)

	// TODO: load resolvers, handle errors
	//var resolvers = make([]*CommitResolver, 0, len(commits))
	//for _, commit := range commits {
	//	resolver := &CommitResolver{Commit: commit}
	//	resolvers = append(resolvers, resolver)
	//}

	var c model.ChangeData
	c.Code = "blas"

	var res []*ChangeResolver
	res[0] = &ChangeResolver{Change: &c}

	return &ChangeConnectionResolver{nil, nil}, nil
}
