package controllers 

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/Tauraih/stockscreener/models"
	"github.com/Tauraih/stockscreener/utils"

	"github.com/dgrijalva/jwt-go"
)

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Signup ...
func Signup(w http.ResponseWriter, r *http.Request){
	var user models.User
	var error models.Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Please enter a valid email"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil{
		log.Println(err)
	}

	user.Password = string(hash)

	models.DB.Create(&user)
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	utils.ResponseJSON(w, user)
}

// Login ...
func Login(w http.ResponseWriter, r *http.Request){
	var user models.User
	var jwt models.JWT
	var error models.Error 

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Please enter a valid email"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password
	models.DB.Where("email = ?", user.Email).First(&user)

	hashedPassword := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil{
		error.Message = "Invalid email or password"
		utils.RespondWithError(w, http.StatusUnauthorized, error)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil{
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	utils.ResponseJSON(w, jwt)
}

// TokenVerifyMiddleware ...
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
				errorObject.Message = "Invalid token"
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
	})
}

