package main

import (
	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// ./context/migration/201807060954_CreateReleaseTable.go
func m201807060954CreateReleaseTable(db *gorm.DB) *gormigrate.Migration {
	m := gormigrate.Migration{
		ID: "201807060954",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&Release{}).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(&Release{}).Error
		},
	}
	return &m
}
