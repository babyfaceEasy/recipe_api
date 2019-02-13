package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/babyfaceeasy/recipe_api/models"

	"github.com/jinzhu/gorm"
)

type test_struct struct {
	Test string
}

type RecipeForm struct {
	Name       string
	PrepTime   string
	Difficulty int
	Vegetarian bool
}

type MyResponse struct {
	Status  int
	Message string
	data    []models.Recipe
}

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

// NewRecipe called when you want to create a new recipe
func NewRecipe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var recipe models.Recipe
	err := decoder.Decode(&recipe)

	log.Println(recipe.PrepTime)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/recipedemo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	// try saving it into the db
	db.Create(&recipe)

	var myResp MyResponse

	myResp.Status = http.StatusCreated
	myResp.Message = "Recipe created successfully"
	myResp.data = nil

	// Marshal or convert the myResp back to Json
	myRespJSON, err := json.Marshal(myResp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type Header so that our clients would know how to read it.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write json back to response
	w.Write(myRespJSON)
}
