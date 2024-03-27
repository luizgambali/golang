package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func CreateTable() {
	//Create a new sqlite database
	db, err := sql.Open("sqlite", "test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		const query string = `
			CREATE TABLE IF NOT EXISTS members (
			id INTEGER NOT NULL PRIMARY KEY,
			name CHAR(40) NOT NULL,
			mtype CHAR(100) NOT NULL,
			email CHAR(50),
			address TEXT NOT NULL,
			mobile CHAR(25) NOT NULL);`
		_, err := db.Exec(query)

		if err != nil {
			//Add the error message to the log
			log.Fatal(err)
		} else {
			//Print the success message
			fmt.Println("Table is created successfully.")
		}

	}
	//Close the database connection
	db.Close()
}
