package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func ConnectDatabase() {

	//Create a new SQLite database
	db, err := sql.Open("sqlite", "test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		//Print the success message
		fmt.Println("Database is connected successfully.")
	}
	//Close the database connection
	db.Close()
}

func Conectar() {

}
