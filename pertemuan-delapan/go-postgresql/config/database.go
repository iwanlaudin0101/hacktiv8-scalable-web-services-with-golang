package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitDb() {
	db, err = sql.Open("postgres", Config())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func GetDb() *sql.DB {
	return db
}
