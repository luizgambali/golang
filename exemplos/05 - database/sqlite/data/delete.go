package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Delete() {
	//Create a new SQLite database
	db, err := sql.Open("sqlite", "test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		//Define the delete query
		const query string = `DELETE FROM members WHERE id = 3;`
		//Execute the query
		_, err := db.Exec(query)

		if err != nil {
			//Add the error message to the log
			log.Fatal(err)
		} else {
			//Print the success message
			fmt.Println("Record is deleted successfully.")
		}
	}
	//Close the database connection
	db.Close()
}
