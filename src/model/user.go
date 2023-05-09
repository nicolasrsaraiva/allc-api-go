package model

import (
	"github.com/nicolasrsaraiva/allc-api/src/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	State    string `json:"state"`
	City     string `json:"city"`
	Street   string `json:"street"`
	District string `json:"district"`
	Number   string `json:"number"`
}

func (user *User) CreateUser() error {
	db := db.ConnectDB()
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO users(Name, Email, Phone, Password, State, City, Street, District, Number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Phone, hashedPassword, user.State, user.City, user.Street, user.District, user.Number)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
