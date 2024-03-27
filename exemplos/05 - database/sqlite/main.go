package main

import (
	"exemplo01.sqlite/data"

	_ "modernc.org/sqlite"
)

func main() {
	// data.CreateDatabase()
	// data.ConnectDatabase()
	// data.CreateTable()
	//data.Insert()
	//data.Update()
	//data.Delete()
	data.Load()
}
