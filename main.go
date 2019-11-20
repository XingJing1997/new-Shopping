package main

import (
	"shopping/config"
	"shopping/db"
	"shopping/router"
)

func main() {
	config.InitConfig()
	db.InitDB()
	router.InitRouter()
}
