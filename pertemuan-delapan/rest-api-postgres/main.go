package main

import (
	"rest-api-postgres/config"
	"rest-api-postgres/routes"
)

func main() {
	routes.StartServerApi().Run(config.APIPORT)
}
