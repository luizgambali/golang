package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Load() {
	//Create a new SQLite database
	db, err := sql.Open("sqlite", "test.db")

	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {

		//Declare variables to store the row values
		var id int
		var name string
		var mtype string
		var address string
		var mobile string
		const query string = `SELECT id, name, mtype, address, mobile
								  FROM members
								  WHERE mtype = 'Silver';`

		//Execute the query
		rows, err := db.Query(query)

		if err != nil {
			//Add the error message to the log
			log.Fatal(err)
		} else {

			//Print the success message
			fmt.Println("Records of all silver members:")
			fmt.Println("ID\tName\t\tMember Type\tAddress\t\tContact No")
			for rows.Next() {
				rows.Scan(&id, &name, &mtype, &address, &mobile)
				fmt.Printf("%d\t %s\t %s\t %s\t %s\n", id, name, mtype, address, mobile)
			}
		}
	}
	//Close the database connection
	db.Close()
}
