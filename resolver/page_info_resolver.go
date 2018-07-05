package resolver

import (
	"changelog/relay"
)

// PageInfoResolver :
type PageInfoResolver struct {
	PageInfo *relay.PageInfo
}

// HasNextPage :
func (r *PageInfoResolver) HasNextPage() bool {
	if r.PageInfo != nil {
		return r.PageInfo.HasNextPage
	}
	return false
}

// HasPreviousPage :
func (r *PageInfoResolver) HasPreviousPage() bool {
	if r.PageInfo != nil {
		return r.PageInfo.HasPreviousPage
	}
	return false
}

// StartCursor :
func (r *PageInfoResolver) StartCursor() string {
	if r.PageInfo != nil {
		return string(r.PageInfo.StartCursor)
	}
	return ""
}

// EndCursor :
func (r *PageInfoResolver) EndCursor() string {
	if r.PageInfo != nil {
		return string(r.PageInfo.EndCursor)
	}
	return ""
}
