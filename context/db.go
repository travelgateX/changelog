package context

import (
	"github.com/jinzhu/gorm"

	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDB : new database connection
func OpenDB(conStr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MustOpenDB : new database connection, panics on error
func MustOpenDB(conStr string) *gorm.DB {
	db, err := OpenDB(conStr)
	if err != nil {
		panic("fatal error connecting database: " + err.Error())
	}
	return db
}
