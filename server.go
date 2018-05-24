package main

import (
	"changelog/context"
	log "changelog/log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Load config, database and set log lvl
	config := context.LoadConfig(".")
	log.SetLevel(config.DebugMode)
	db, err = context.OpenDB(config)

	// AutoMigrate : Check its working ok!! <=================
	context.AutoMigrate(db)

	// Config url routes
	router := mux.NewRouter()
	router.HandleFunc("/changes", GetChanges).Methods("GET")

	// Start HTTP server
	log.Info("== Changelog server ON == Listening at port:%s ==", config.HTTPPort)
	if err = http.ListenAndServe(":"+config.HTTPPort, router); err != nil {
		log.Fatal("Changelog server has exploited: %s", err)
	}

}

// GetChanges : All changes handler
func GetChanges(w http.ResponseWriter, r *http.Request) {
	var changes []context.Change
	db.Find(&changes)
	json.NewEncoder(w).Encode(&changes)
}
