package main

import (
	"assignment-2/database"
	"assignment-2/routers"

	_ "github.com/lib/pq"
)

func main() {
	database.StartDB()

	var PORT = ":8123"

	routers.StartServer().Run(PORT)
}
