package main

import (
	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// ./context/migration/2201807041524_CreateSourceTable.go
func m201807041524CreateSourceTable(db *gorm.DB) *gormigrate.Migration {
	m := gormigrate.Migration{
		ID: "201807041524",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&Source{}).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(&Source{}).Error
		},
	}
	return &m
}
