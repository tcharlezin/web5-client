package controllers

import (
	"client/models"
	"client/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	services.CreateRSAKey()
	json.NewEncoder(w).Encode(services.ReadRSAPublicKey())
}

func Login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.ReadRSAPublicKey())
}
