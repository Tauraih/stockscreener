package utils

import (
	"encoding/json"
	"encoding/csv"
	"net/http"
	"log"
	"os"
	"time"

	"github.com/Tauraih/stockscreener/models"
	"github.com/dgrijalva/jwt-go"
	)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, status int, error models.Error){
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(error)
}


// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}


// GenerateToken ...
func GenerateToken(user models.User) (string, error){
	var err error

	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
		"issuer": "portfolio",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil{
		log.Println(err)
	}

	// spew.Dump(tokenString)
	return tokenString, nil
}


//ReadCsv accepts a file and returns its content as a multi-dimensional type 
//with lines and each column, Only parses to string type
func ReadCsv(fileName string) ([][]string, error) {
	// open csv file 
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close() 

	r := csv.NewReader(f)
	r.Comma = ';'

	lines, err := r.ReadAll()
	// lines.Comma = ';'
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
} 


// Remove $ sign from 
func RemoveDollar(field string) (string, string){
	return "", ""
}

