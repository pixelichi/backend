package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// This function will make a connection to the database only once.
func init() {
	var err error

	connStr := "postgres://admin:password@localhost/4me?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("This is main!")
	db.Ping()

	// insert
	// hardcoded
	insertStmt := `insert into "dog"("first_name", "last_name", "color", "chip_number") values('nibbler', 'ta', 'black', 421)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jack", 21)
	// CheckError(e)
}
