package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

func TestPGConnect() {
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, user)
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE " + "abcd")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("drop DATABASE " + "abcd")
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
