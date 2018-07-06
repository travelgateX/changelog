package model

import (
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// Change :
type Change struct {
	Code       graphql.ID
	ChangeData *ChangeData
	CreatedAt  datetime.DateTime
	UpdatedAt  datetime.DateTime
}

// ChangeData :
type ChangeData struct {
	Code       graphql.ID
	Message    string
	Project    *string
	ChangeDate datetime.DateTime
	Type       ChangeType
	Release    Release
	Source     Source
	CreatedAt  datetime.DateTime
	UpdatedAt  datetime.DateTime
}

// ChangeType : Change type
type ChangeType string

// Change type enum
const (
	Added      = ChangeType("ADDED")
	Changed    = ChangeType("CHANGED")
	Deprecated = ChangeType("DEPRECATED")
	Removed    = ChangeType("REMOVED")
	Fixed      = ChangeType("FIXED")
	Security   = ChangeType("SECURITY")
)
