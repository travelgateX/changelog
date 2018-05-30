package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDB : OpenDB
func OpenDB(config *Config) (*gorm.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		return nil, errors.New("error connecting database using [%s]", conStr)
	}

	// Debug mode
	db.LogMode(config.SQLDebugMode)

	return db, nil
}

// MigrateFullDB : migrate all DB schema
func MigrateFullDB(db *gorm.DB) *gorm.DB {
	db = db.AutoMigrate(&ChangeType{}, &Resource{}, &Release{}, &Change{})
	db.Model(&Release{}).AddForeignKey("resource_id", "resources(id)", "RESTRICT", "RESTRICT")
	db.Model(&Change{}).AddForeignKey("release_id", "releases(id)", "RESTRICT", "RESTRICT")

	return db
}

// ChangeType : model
type ChangeType struct {
	gorm.Model
	Name    string
	Changes []Change
}

// Resource : model
type Resource struct {
	gorm.Model
	Name     string
	Releases []Release
}

// Release : model
type Release struct {
	gorm.Model
	X, Y, Z    int
	Delivery   time.Time
	ResourceID int
	Changes    []Change
}

// Change : model
type Change struct {
	gorm.Model
	Body         string
	ChangeTypeID int
	ReleaseID    int
}
