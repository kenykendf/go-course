package main

import (
	"final-project/database"
	"final-project/router"
)

func main() {
	r := router.Init()

	database.StartDB()

	r.Run(":8085")
}
