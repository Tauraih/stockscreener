package controllers 

import (
	"fmt"
	"encoding/json"
	"encoding/csv"
	"net/http"
	"strings"

	"github.com/Tauraih/stockscreener/models"
	"github.com/Tauraih/stockscreener/utils"

	"github.com/gorilla/mux"  
)

type upload struct {
	Content []byte
}

// GetBusinesses ...
func GetBusinesses(w http.ResponseWriter, r *http.Request) {
	var companies []models.New
	// models.DB.Find(&companies)
	models.DB.Distinct("name", "symbol").Order("name desc").Find(&companies)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, companies)
}

// GetBusiness ... 
func GetBusiness(w http.ResponseWriter, r *http.Request) {
	var company models.New
	params := mux.Vars(r)
	models.DB.First(&company, params["id"])

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// GetByBusiness ...
func GetByBusiness(w http.ResponseWriter, r *http.Request) {
	var company []models.New
	params := mux.Vars(r)
	models.DB.Where("symbol = ?", params["symbol"]).Find(&company)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// CreateRecord ...
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var company models.New
	json.NewDecoder(r.Body).Decode(&company)
	models.DB.Create(&company)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// CreateFromCsv ...
func CreateFromCsv(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	rs := csv.NewReader(f)
	rs.Comma = ';'

	lines, err := rs.ReadAll()
	
	if err != nil {
		fmt.Println(err)
	}

	for i, line := range lines {
		data := models.New{
			Number: line[0],
			Name:  line[1],
			Symbol:  line[2],
			High:  strings.Split(line[3], " ")[1],
			Low:   strings.Split(line[4], " ")[1],
			Opening: strings.Split(line[5], " ")[1],
			Last:  strings.Split(line[6], " ")[1],
			Closing: strings.Split(line[7], " ")[1],
			Change: strings.Replace(line[8], "%", "", -1),
		}

		if i != 0 {
			models.DB.Create(&data)
			fmt.Println(data.Change + "   	" + data.Low + "		" + data.Last + "		" + data.Opening + "		" + data.Closing)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, "business")
}

// UpdateRecord ... 
func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	var company models.New
	var error models.Error

	params := mux.Vars(r)
	models.DB.First(&company, params["id"])
	json.NewDecoder(r.Body).Decode(&company)

	if company.Name == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Symbol == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.High == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Low == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Opening == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Closing == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Last == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if company.Change == "" {
		error.Message = "This field can not be empty"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	models.DB.Save(&company)

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, company)
}

// DeleteRecord ... 
func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var company models.Company 
	models.DB.Delete(&company, params["id"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("company deleted successfully")
}
