package main

import (
	"log"
	"net/http"

	"github.com/Tauraih/stockscreener/controllers"
	"github.com/Tauraih/stockscreener/models"

	"github.com/gorilla/mux"
)

func initializeRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.TokenVerifyMiddleware(controllers.GetUsers)).Methods("GET")
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/companies", controllers.GetCompanies).Methods("GET")
	r.HandleFunc("/company/{id}", controllers.GetCompany).Methods("GET")
	r.HandleFunc("/company", controllers.CreateCompany).Methods("POST")
	r.HandleFunc("/company/{id}", controllers.UpdateCompany).Methods("PATCH")
	r.HandleFunc("/company/{id}", controllers.DeleteCompany).Methods("PATCH")

	r.HandleFunc("/prices", controllers.GetPrices).Methods("GET")
	r.HandleFunc("/prices/{id}", controllers.GetPrice).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	models.InitialMigration()
	initializeRouter()
}
