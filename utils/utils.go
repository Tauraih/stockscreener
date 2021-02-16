package utils

import (
	"encoding/json"
	"net/http"
	"log"
	"os"

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
		"issuer": "portfolio",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil{
		log.Println(err)
	}

	// spew.Dump(tokenString)
	return tokenString, nil
}

