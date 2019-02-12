package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// HomeHandler responds to /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	jsonVal, err := json.Marshal(&struct{ name string }{name: "olakunle"})

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonVal)
}
