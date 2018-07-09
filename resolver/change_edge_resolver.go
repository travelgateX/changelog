package resolver

import (
	"changelog/model"
	"changelog/relay"

	"github.com/jinzhu/gorm"
)

// ChangeEdgeResolver :
type ChangeEdgeResolver struct {
	db   *gorm.DB
	edge *relay.Edge
	err  *[]*model.Error
}

// Cursor :
func (r *ChangeEdgeResolver) Cursor() string {
	if r.edge != nil {
		return string(r.edge.Cursor)
	}
	return ""
}

// Node :
func (r *ChangeEdgeResolver) Node() *ChangeResolver {
	if r.err == nil {
		x := r.edge.Node.(*model.ChangeData)
		return &ChangeResolver{change: x, err: nil}
	}
	return &ChangeResolver{change: nil, err: r.err}
}
