package relay

// Filters :
type Filters struct {
	First  *int32
	Last   *int32
	Before *string
	After  *string
	Codes  *[]string
}

// Filters :
func (r *Filters) Filters() map[string]interface{} {
	filters := make(map[string]interface{})
	if r.First != nil {
		filters["first"] = int(*r.First)
	}
	if r.Last != nil {
		filters["last"] = int(*r.Last)
	}
	if r.Before != nil {
		filters["before"] = *r.Before
	}
	if r.After != nil {
		filters["after"] = *r.After
	}
	return filters
}
