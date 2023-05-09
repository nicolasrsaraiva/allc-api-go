package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	nameDB := os.Getenv("NAME_DB")
	passDB := os.Getenv("PASS_DB")

	conn := fmt.Sprintf("user=postgres dbname=%s password=%s host=localhost sslmode=disable", nameDB, passDB)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
