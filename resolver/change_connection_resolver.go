package resolver

import (
	"changelog/model"
	"changelog/relay"

	"github.com/jinzhu/gorm"
)

// ChangeConnectionResolver :
type ChangeConnectionResolver struct {
	db   *gorm.DB
	conn *relay.Connection
	err  *[]*model.Error
}

// PageInfo :
func (r *ChangeConnectionResolver) PageInfo() *PageInfoResolver {
	if r.err == nil {
		return &PageInfoResolver{&r.conn.PageInfo}
	}
	return &PageInfoResolver{nil}
}

// Edges :
func (r *ChangeConnectionResolver) Edges() *[]*ChangeEdgeResolver {
	if r.err != nil {
		resolvers := make([]*ChangeEdgeResolver, 0, 1)
		resolvers = append(resolvers, &ChangeEdgeResolver{db: r.db, edge: nil, err: r.err})
		return &resolvers
	}

	resolvers := make([]*ChangeEdgeResolver, 0, len(r.conn.Edges))
	for i := range r.conn.Edges {
		resolvers = append(resolvers, &ChangeEdgeResolver{db: r.db, edge: r.conn.Edges[i], err: nil})
	}
	return &resolvers
}
