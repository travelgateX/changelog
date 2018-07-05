package model

import (
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// Change :
type Change struct {
	Code       graphql.ID `db:"code_id" gorm:"type:serial"`
	ChangeData *ChangeData
	CreatedAt  datetime.DateTime `db:"created_at"`
	UpdatedAt  datetime.DateTime `db:"updated_at"`
}

// ChangeData :
type ChangeData struct {
	Code       graphql.ID        `db:"code_id" gorm:"type:serial"`
	Message    string            `db:"message"`
	Project    *string           `db:"project"`
	ChangeDate datetime.DateTime `db:"change_date"`
	Type       ChangeType        `db:"change_type"`
	Release    Release
	Source     Source
	CreatedAt  datetime.DateTime `db:"created_at"`
	UpdatedAt  datetime.DateTime `db:"updated_at"`
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
