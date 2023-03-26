package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func CreateConnection() *sql.DB {
	db, err = sql.Open("postgres", ConnectionString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
