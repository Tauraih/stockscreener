package controllers 

import ( 
	// "encoding/json"
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
