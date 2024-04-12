package main

import (
	"github.com/luizgambali/crud-api-1/db"
	"github.com/luizgambali/crud-api-1/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
