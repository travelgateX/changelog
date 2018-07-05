package resolver

import (
	"changelog/model"
	"changelog/relay"
)

// ChangeEdgeResolver :
type ChangeEdgeResolver struct {
	Edge  *relay.Edge
	Error *[]*model.Error
}

// Cursor :
func (r *ChangeEdgeResolver) Cursor() string {
	if r.Edge != nil {
		return string(r.Edge.Cursor)
	}
	return ""
}

// Node :
func (r *ChangeEdgeResolver) Node() *ChangeResolver {
	var c model.ChangeData
	c.Code = "blas"

	res := &ChangeResolver{Change: &c}

	return res

	// if r.Error == nil {
	// 	x := r.Edge.Node.(*api.APIData)
	// 	return &ChangeResolver{nil, x, nil, r.Permission}
	// } else {
	// 	return &ChangeResolver{nil, nil, r.Error, nil}
	// }
}
