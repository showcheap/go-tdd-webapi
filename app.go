package main

import (
	"encoding/json"
	"go-api-tdd/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// App ...
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize ...
func (a *App) Initialize(dbname string) {

	// Init Router
	a.Router = mux.NewRouter()
	a.InitializeRoutes()

	// Init DB
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("Failed to connect db")
	}

	db.AutoMigrate(&models.Product{})

	a.DB = db
}

// InitializeRoutes ...
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/product", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
}

// Run ...
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// Error return error response
func (a *App) Error(w http.ResponseWriter, code int, message string) {
	a.JSON(w, code, map[string]string{"error": message})
}

// JSON return josn response
func (a *App) JSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	q := a.DB.Find(&products)

	if q.Error != nil {
		a.Error(w, http.StatusInternalServerError, q.Error.Error())
		return
	}

	a.JSON(w, http.StatusOK, products)
	return
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var product models.Product

	q := a.DB.Where("id = ?", vars["id"]).First(&product)

	if q.RecordNotFound() {
		a.Error(w, http.StatusNotFound, "Product not found")
		return
	}

	a.JSON(w, http.StatusOK, product)
	return
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		a.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	q := a.DB.Create(&product)

	if q.Error != nil {
		a.Error(w, http.StatusInternalServerError, q.Error.Error())
		return
	}

	a.JSON(w, http.StatusCreated, product)

}
