package main

import (
	"changelog/context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Load config and database
	config := context.LoadConfig(".")
	db, err = context.OpenDB(config)

	// AutoMigrate : Check its working ok!! <=================
	context.AutoMigrate(db)

	// Config url routes
	router := mux.NewRouter()
	router.HandleFunc("/changes", GetChanges).Methods("GET")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":"+config.HTTPPort, router))
}

// GetChanges : All changes handler
func GetChanges(w http.ResponseWriter, r *http.Request) {
	var changes []context.Change
	db.Find(&changes)
	json.NewEncoder(w).Encode(&changes)
}
