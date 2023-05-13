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

func GetUsers() []User {
	users := []User{}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	stmt, err := dbConn.Query("SELECT Id, Name, Email, Phone, Password, State, City, Street, District, Number FROM users ORDER BY ID")

	if err != nil {
		panic(err.Error())
	}

	var (
		id, name, email, phone, password, state, city, street, district, number string
	)

	for stmt.Next() {

		err := stmt.Scan(&id, &name, &email, &phone, &password, &state, &city, &street, &district, &number)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, User{id, name, email, phone, password, state, city, street, district, number})
	}

	defer stmt.Close()

	return users
}

func GetUserByEmail(userEmail string) User {
	dbConn := db.ConnectDB()
	defer dbConn.Close()

	var user User

	stmt, err := dbConn.Prepare("SELECT * FROM users WHERE email=$1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(userEmail).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, &user.State, &user.City, &user.Street, &user.District, &user.Number)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func (user *User) GetUser() error {
	return nil
}
