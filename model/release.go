package model

import (
	"changelog/scalars/datetime"
	"changelog/scalars/version"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Release :
type Release struct {
	ReleaseData *ReleaseData
}

// ReleaseData :
type ReleaseData struct {
	gorm.Model
	Code        graphql.ID        `gorm:"type:varchar(100)"`
	Name        string            `gorm:"type:varchar(255)"`
	Version     version.Version   `gorm:"type:varchar(255)"`
	releaseDate datetime.DateTime `gorm:"type:timestamp"`
}

// TableName :
func (ReleaseData) TableName() string {
	return "releases"
}
