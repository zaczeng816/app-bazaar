package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"app-bazaar/model"
	"app-bazaar/service"

	jwt "github.com/form3tech-oss/jwt-go"
)

var mySigningKey = []byte("secret")

func signinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a signup request")
	w.Header().Set("Content-Type", "text/plain")
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Failed to decode user %v\n", err)
		return
	}

	exists, err := service.CheckUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to read user from database", http.StatusInternalServerError)
		fmt.Printf("Failed to read user from database %v\n", err)
		return
	}
	if !exists {
		http.Error(w, "User does not exist or password is incorrect", http.StatusUnauthorized)
		fmt.Println("User does not exist")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Failed to generate token %v\n", err)
		return
	}

	w.Write([]byte(tokenString))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a signup request")
	w.Header().Set("Content-Type", "text/plain")
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Failed to decode user %v\n", err)
		return
	}

	if user.Username == "" || len(user.Password) < 6 ||
		!regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(user.Username){
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password: %v, %v", user.Username, user.Password)
		return
	}

	success, err := service.AddUser(&user)
	if err != nil {
		http.Error(w, "Failed to save user to database", http.StatusInternalServerError)
		fmt.Printf("Failed to save user to database %v\n", err)
		return
	}

	if !success {
		http.Error(w, "User already exists", http.StatusBadRequest)
		fmt.Println("User already exists")
		return
	}

	fmt.Printf("User %s successfully created\n", user.Username)
}