package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// HomeHandler this responds to the / endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla\n"))
}

func main() {
	// orm library setup
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/recipedemo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to open")
	}
	log.Println("Connection Established")

	r := mux.NewRouter()

	// add endpoints here
	r.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
