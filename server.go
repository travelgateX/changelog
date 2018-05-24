package main

import (
	"changelog/context"
	log "changelog/log"
	"changelog/tools"
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

	// MigrateFullDB : Check its working ok!! <=================
	context.MigrateFullDB(db)

	// Config url routes
	router := mux.NewRouter()
	router.HandleFunc("/changes", GetChanges).Methods("GET")
	router.HandleFunc("/change/{id}", GetChange).Methods("GET")
	router.HandleFunc("/change", CreateChange).Methods("POST")

	// Start HTTP server
	log.Info("== Changelog server ON == Listening at port:%s ==", config.HTTPPort)
	if err = http.ListenAndServe(":"+config.HTTPPort, router); err != nil {
		log.Fatal("Changelog server has exploited: %s", err)
	}

}

// GetChanges : [GET] Changes List
func GetChanges(w http.ResponseWriter, r *http.Request) {
	logTransaction(r)

	var changes []context.Change
	db.Find(&changes)
	json.NewEncoder(w).Encode(&changes)
}

// GetChange : [GET] Change by ID
func GetChange(w http.ResponseWriter, r *http.Request) {
	logTransaction(r)

	params := mux.Vars(r)
	var change context.Change
	db.First(&change, params["id"])
	json.NewEncoder(w).Encode(&change)
}

// CreateChange : [POST] Create Change
func CreateChange(w http.ResponseWriter, r *http.Request) {
	logTransaction(r)

	var change context.Change
	json.NewDecoder(r.Body).Decode(&change)
	db.Create(&change)
	json.NewEncoder(w).Encode(&change)
}

// Log transaction with params
func logTransaction(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Error("Parse form has exploited: %s", err)
	}
	log.Info("[API REQUEST] [%s] [%s]", tools.Path(r), r.Form)
}
