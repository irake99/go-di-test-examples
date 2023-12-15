package di

import (
	"database/sql"
	"fmt"
	"log"
)

func queryWithoutDI() {
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

// ------------------------------------------------------------------------

const dbConnectionString = "username:password@tcp(localhost:3306)/dbname"

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func queryWithDI(db Database) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

func Test1() {
	queryWithoutDI()

	// ----------------------------------------------------

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryWithDI(db)
}
