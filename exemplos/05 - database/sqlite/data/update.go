package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Update() {
	//Create a new SQLite database
	db, err := sql.Open("sqlite", "test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		const query string = `
			UPDATE members SET mobile = '018563452390' WHERE id = 2;`
		_, err := db.Exec(query)

		if err != nil {
			//Add the error message to the log
			log.Fatal(err)
		} else {
			//Print the success message
			fmt.Println("Record is updated successfully.")
		}
	}
	//Close the database connection
	db.Close()
}
