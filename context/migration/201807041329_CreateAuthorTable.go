package main

import (
	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// ./context/migration/201806051541_CreateCommitsTable.go
func m201807041329CreateAuthorTable(db *gorm.DB) *gormigrate.Migration {
	m := gormigrate.Migration{
		ID: "201807041329",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&Author{}).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(&Author{}).Error
		},
	}
	return &m
}
