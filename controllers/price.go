package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Tauraih/stockscreener/models"
	"github.com/Tauraih/stockscreener/utils"

	"github.com/gorilla/mux"
)

// GetPrices ...
func GetPrices(w http.ResponseWriter, r *http.Request) {
	var prices []models.Prices
	models.DB.Find(&prices)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, prices)
}

// GetPrice ... 
func GetPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var price models.Prices
	models.DB.First(&price, params["id"])

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, price)
}

// AddPrice ... 
func AddPrice(w http.ResponseWriter, r *http.Request) {
	var price models.Prices
	json.NewDecoder(r.Body).Decode(&price)
    models.DB.Create(&price)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, price)
}

// UpdatePrice ... 
func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	var price models.Prices
	params := mux.Vars(r)
	models.DB.First(&price, params["id"])
	json.NewDecoder(r.Body).Decode(&price)
	models.DB.Save(&price)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, price)
}

// DeletePrice ... 
func DeletePrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var price models.Prices
	models.DB.Delete(&price, params["id"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("company deleted successfully")
}
