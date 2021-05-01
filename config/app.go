package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "khdoba"
	password = "khdoba19012000"
	dbname   = "postgres"
)

func Connect() *sql.DB {
	// Please define your user name and password for my sql.
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
