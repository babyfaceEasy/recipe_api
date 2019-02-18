package main

import (
	"log"
	"net/http"
	"time"

	controller "github.com/babyfaceeasy/recipe_api/controllers"

	"github.com/babyfaceeasy/recipe_api/models"

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

	// Migrate the Schema
	db.AutoMigrate(&models.Recipe{})

	log.Println("Tables created!")

	r := mux.NewRouter()

	// add endpoints here
	r.HandleFunc("/", controller.HomeHandler).Methods("GET").Name("home")
	r.HandleFunc("/recipes", controller.NewRecipe).Methods("POST").Name("newRecipe")
	r.HandleFunc("/recipes", controller.ListRecipes).Methods("GET").Name("listRecipes")
	r.HandleFunc("/recipes/{recipeID}", controller.GetRecipe).Methods("GET").Name("getRecipe")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
