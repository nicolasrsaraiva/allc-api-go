package model

import (
	"github.com/nicolasrsaraiva/allc-api/src/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Phone    string
	Password string
	City     string
	Street   string
}

func (user *User) CreateUser() error {
	db := db.ConnectDB()
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO users(Name, Email, Phone, Password, City, Street) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Phone, hashedPassword, user.City, user.Street)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
