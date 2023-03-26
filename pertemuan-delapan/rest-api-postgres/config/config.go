package config

import "fmt"

const (
	// api config
	APIPORT = ":4000"

	// web config
	WEBPORT = ":5001"

	// database config
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "go-learn-postgres"
)

func ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
