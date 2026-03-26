package DB

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func ConnectDb() *sql.DB {
	var err error

	godotenv.Load("../.env")

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		user, dbname, password)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return Db
}
