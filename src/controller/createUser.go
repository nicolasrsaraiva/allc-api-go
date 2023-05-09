package controller

import (
	"net/http"

	"github.com/nicolasrsaraiva/allc-api/src/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")

	name := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	city := r.FormValue("city")
	street := r.FormValue("street")

	user := model.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		City:     city,
		Street:   street,
	}

	err := user.CreateUser()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
