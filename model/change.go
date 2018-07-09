package model

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Change :
type Change struct {
	ChangeData *ChangeData
}

// ChangeData :
type ChangeData struct {
	gorm.Model
	Code       graphql.ID
	Message    string
	Project    *string
	ChangeDate time.Time
	Type       ChangeType
	ReleaseID  uint
	SourceID   uint
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

// TableName :
func (ChangeData) TableName() string {
	return "changes"
}

// ChangeObject :
func (r ChangeData) ChangeObject() *Change {
	return &Change{ChangeData: &r}
}
