package main

import (
	"time"

	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// ./context/migration/201806051541_CreateCommitsTable.go
func m201806051541CreateCommitsTable(db *gorm.DB) *gormigrate.Migration {
	type Category string
	type Commit struct {
		ID        int32     `json:"id" gorm:"type:serial"`
		Message   string    `json:"message" gorm:"type:text"`
		User      string    `json:"user"`
		Release   string    `json:"release"`
		Resource  string    `json:"resource"`
		Category  Category  `json:"category" gorm:"type:varchar(10)"`
		Released  bool      `json:"released"`
		CreatedAt time.Time `json:"created_at"`
	}

	m := gormigrate.Migration{
		ID: "201806051541",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&Commit{}).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(&Commit{}).Error
		},
	}
	return &m
}
