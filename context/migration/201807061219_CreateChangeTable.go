package main

import (
	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// ./context/migration/201807061219_CreateChangeTable.go
func m201807061219CreateChangeTable(db *gorm.DB) *gormigrate.Migration {
	m := gormigrate.Migration{
		ID: "201807061219",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&Change{}).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(&Change{}).Error
		},
	}
	return &m
}
