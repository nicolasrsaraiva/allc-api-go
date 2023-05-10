package model

import (
	"github.com/nicolasrsaraiva/allc-api/src/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	State    string `json:"state" validate:"required"`
	City     string `json:"city" validate:"required"`
	Street   string `json:"street" validate:"required"`
	District string `json:"district" validate:"required"`
	Number   string `json:"number" validate:"required"`
}

func (user *User) CreateUser() error {
	dbConn := db.ConnectDB()
	defer dbConn.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := dbConn.Prepare("INSERT INTO users(Name, Email, Phone, Password, State, City, Street, District, Number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
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

func (user *User) GetUsers() error {

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	stmt, err := dbConn.Query("SELECT(Name, Email, Phone, State, City, Street, District, Number) FROM users ORDER BY ID")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	return nil
}
