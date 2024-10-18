package main

import (
	"BasicTrade-API/database"
	"BasicTrade-API/router"
)

var (
	PORT = ":9000"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
