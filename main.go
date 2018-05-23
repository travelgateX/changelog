package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
	}
}
