package resolver

import (
	"changelog/model"
	"changelog/relay"
)

// ChangeConnectionResolver :
type ChangeConnectionResolver struct {
	Connection *relay.Connection
	Error      *[]*model.Error
}

// PageInfo :
func (r *ChangeConnectionResolver) PageInfo() *PageInfoResolver {
	if r.Error == nil {
		return &PageInfoResolver{&r.Connection.PageInfo}
	}
	return &PageInfoResolver{nil}
}

// Edges :
func (r *ChangeConnectionResolver) Edges() *[]*ChangeEdgeResolver {

	if r.Error != nil {
		resolvers := make([]*ChangeEdgeResolver, 0, 1)
		resolvers = append(resolvers, &ChangeEdgeResolver{nil, r.Error})
		return &resolvers
	}

	resolvers := make([]*ChangeEdgeResolver, 0, len(r.Connection.Edges))
	for i := range r.Connection.Edges {
		resolvers = append(resolvers, &ChangeEdgeResolver{r.Connection.Edges[i], nil})
	}
	return &resolvers
}
