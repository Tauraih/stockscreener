package controllers 

import (
	"encoding/json"
	"net/http"

	"github.com/Tauraih/stockscreener/models"
	"github.com/Tauraih/stockscreener/utils"

	"github.com/gorilla/mux"  
)

// GetCompanies ...
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	var companies []models.Company 
	models.DB.Find(&companies)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, companies)
}

// GetCompany ... 
func GetCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company 
	params := mux.Vars(r)
	models.DB.First(&company, params["id"])

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// CreateCompany ...
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	json.NewDecoder(r.Body).Decode(&company)
	models.DB.Create(&company)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// UpdateCompany ... 
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	params := mux.Vars(r)
	models.DB.First(&company, params["id"])
	json.NewDecoder(r.Body).Decode(&company)
	models.DB.Save(&company)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// DeleteCompany ... 
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var company models.Company 
	models.DB.Delete(&company, params["id"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("company deleted successfully")
}
