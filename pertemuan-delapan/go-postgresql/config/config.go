package config

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "go-learn-postgres"
)

func Config() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
