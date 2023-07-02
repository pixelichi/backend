package db

import (
	"database/sql"
	"log"

	db "shinypothos.com/db/sqlc"
)

func GetDBStore(dbDriver string, dbSource string) *db.Store {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	// New connection to DB
	db := db.NewStore(conn)

  return db
}
