package context

import (
	"changelog/config"

	"github.com/jinzhu/gorm"

	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDB : new database connection
func OpenDB(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", c.GetDBConnString())
	if err != nil {
		return nil, err
	}

	// DB debug mode
	db.LogMode(c.DebugMode)
	return db, nil
}

// MustOpenDB : new database connection, panics on error
func MustOpenDB(c *config.Config) *gorm.DB {
	db, err := OpenDB(c)
	if err != nil {
		panic("fatal error connecting database: " + err.Error())
	}
	return db
}
