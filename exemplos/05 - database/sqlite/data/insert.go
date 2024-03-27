package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Insert() {

	//Create a new SQLite database
	db, err := sql.Open("sqlite", "test.db")
	//Check for any error
	if err != nil {
		//Add the error message to the log
		log.Fatal(err)
	} else {
		const query string = `
                INSERT INTO members (id, name, mtype, email, address, mobile)
                VALUES(1, 'Nehal Ahmed', 'Silver', 'nehal@gmail.com','36, Dhanmondi 2, Dhaka','01844657342'),
                      (2, 'Abir Chowdhury', 'Gold', 'abir@gmail.com','102, Mirpur 10, Dhaka','01994563423'),
                      (3, 'Mirza Abbas', 'Silver', 'abbas@gmail.com','12, Jigatala, Dhaka','01640006710');`

		//Execute the query
		_, err := db.Exec(query)

		if err != nil {
			//Add the error message to the log
			log.Fatal(err)
		} else {
			//Print the success message
			fmt.Println("Records inserted successfully.")
		}
	}

	//Close the database connection
	db.Close()
}
