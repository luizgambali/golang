package data

import (
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func CreateDatabase() {

	//Create a new SQLite database
	db, err := os.Create("test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		//Print the success message
		fmt.Println("Database is created.")
	}
	//Close the database connection
	db.Close()
}
