package main

import (
	"changelog/config"
	"changelog/context"
	"log"
	"os"

	gorm "github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

func main() {

	var (
		err error
		c   *config.Config
		db  *gorm.DB
	)

	// Load config file
	if c, err = config.LoadConfig("./config"); err != nil {
		log.Fatalf("fatal error, cant parse config file. %v", err)
	}

	// Open database connection
	if db, err = context.OpenDB(c); err != nil {
		log.Fatalf("fatal error, cant open database connection. %v", err)
	}
	defer db.Close()
	db.LogMode(true)

	// Fetch all migrations
	allMigrations := []*gormigrate.Migration{
		m201806051541CreateCommitsTable(db),
	}

	// Migrate all
	m := gormigrate.New(db, gormigrate.DefaultOptions, allMigrations)

	// Options available
	// - migrate: run all migrations *default
	// - rollback: rollback last migration
	option := "migrate"
	if len(os.Args[:1]) > 0 {
		option = os.Args[1]
	}

	switch option {
	case "migrate":
		if err = m.Migrate(); err != nil {
			log.Fatalf("couldnt migrate all migrations:%v", err)
		}
	case "rollback":
		if err = m.RollbackLast(); err != nil {
			log.Fatalf("couldnt rollback last migration:%v", err)
		}
	}

	log.Printf("migration finished successfully")
}
