package context

import (
	log "changelog/log"
	"fmt"

	"github.com/jinzhu/gorm"

	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDB : OpenDB wrap
func OpenDB(config *Config) (*gorm.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		log.Fatal("Fatal error connecting database: %s \n", err)
	}
	return db, nil
}

// AutoMigrate : AutoMigrate wrap
func AutoMigrate(db *gorm.DB) *gorm.DB {
	return db.AutoMigrate(&Change{})
}
