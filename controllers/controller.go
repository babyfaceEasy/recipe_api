package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/babyfaceeasy/recipe_api/models"
	"github.com/jinzhu/gorm"

	valid "github.com/asaskevich/govalidator"
)

type RecipeForm struct {
	Name       string `json:"name" valid:"required"`
	PrepTime   string `json:"prepTime" valid:"required"`
	Difficulty int    `json:"difficulty" valid:"required"`
	Vegetarian bool   `json:"vegetarian" valid:"bool,optional"`
}

type MyResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []models.Recipe `json:"data"`
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
	valid.SetNilPtrAllowedByRequired(true)
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

	//var recipe models.Recipe
	var recipe RecipeForm
	err := decoder.Decode(&recipe)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = valid.ValidateStruct(recipe)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipeModel := models.Recipe{
		Name:       recipe.Name,
		Difficulty: recipe.Difficulty,
		PrepTime:   recipe.PrepTime,
		Vegetarian: recipe.Vegetarian,
	}

	// connecting to my db
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/recipedemo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//save to db
	db.Create(&recipeModel)

	// respond to the world
	myResp := MyResponse{
		Status:  http.StatusCreated,
		Message: "Recipe created successfully!",
		Data:    nil,
	}

	// convert my response back to json
	myRespJSON, err := json.Marshal(myResp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set headers
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write response back
	w.Write(myRespJSON)
}

// ListRecipes returns all the recipes in our database.
func ListRecipes(w http.ResponseWriter, r *http.Request) {
	//connect to the db
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/recipedemo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// var to hold all recipes
	var recipes []models.Recipe

	// get all recipes back
	db.Find(&recipes)

	// create ur response
	myResp := MyResponse{
		Status:  http.StatusOK,
		Message: "List of recipes in the database",
		Data:    recipes,
	}

	// create json out of recipes
	myRespJSON, err := json.Marshal(myResp)

	if err != nil {
		log.Println("Couldn't marshal it")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//set header values
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write to output
	w.Write(myRespJSON)
}

// GetRecipe return all the information relating to a recipe, based on its ID
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	// get the variables passed
	vars := mux.Vars(r)
	recipeID, err := strconv.Atoi(vars["recipeID"])

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// connect to db
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/recipedemo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recipe models.Recipe
	// get recipe based on id
	db.First(&recipe, recipeID)

	if (models.Recipe{}) == recipe {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// create response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write back to browser
	json.NewEncoder(w).Encode(recipe)
}
