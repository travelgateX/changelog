package model

import (
	"github.com/jinzhu/gorm"

	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : database object struct
type DB struct {
	db *gorm.DB
}

// MustOpenDB : new database connection ...
func MustOpenDB(conStr string) *gorm.DB {
	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		panic("fatal error connecting database: " + err.Error())
	}
	return db
}
