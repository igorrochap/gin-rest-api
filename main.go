package main

import (
	"github.com/igorrochap/gin-rest-api/database"
	"github.com/igorrochap/gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
