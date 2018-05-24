package context

import (
	"time"

	"github.com/jinzhu/gorm"
)

// MigrateFullDB : migrate all DB schema
func MigrateFullDB(db *gorm.DB) *gorm.DB {
	db = db.AutoMigrate(&ChangeType{}, &Resource{}, &Release{}, &Change{})

	db.Model(&Release{}).AddForeignKey("resource_id", "resources(id)", "RESTRICT", "RESTRICT")
	db.Model(&Change{}).AddForeignKey("release_id", "releases(id)", "RESTRICT", "RESTRICT")

	return db
}

// ChangeType : model
type ChangeType struct {
	ID      int
	Name    string
	Changes []Change
}

// Resource : model
type Resource struct {
	ID       int
	Name     string
	Releases []Release
}

// Release : model
type Release struct {
	gorm.Model
	X, Y, Z    int
	Date       time.Time
	ResourceID int
	Changes    []Change
}

// Change : model
type Change struct {
	gorm.Model
	Body         string
	Date         time.Time
	ChangeTypeID int
	ReleaseID    int
}
