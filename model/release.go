package model

import (
	"changelog/scalars/version"
	"time"

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
	Code        graphql.ID      `gorm:"type:varchar(100)"`
	Name        string          `gorm:"type:varchar(255)"`
	Version     version.Version `gorm:"type:varchar(255)"`
	ReleaseDate time.Time       `gorm:"type:timestamp"`
	Changes     []Change
}

// TableName :
func (ReleaseData) TableName() string {
	return "releases"
}

// ReleaseObject :
func (r ReleaseData) ReleaseObject() *Release {
	return &Release{ReleaseData: &r}
}
