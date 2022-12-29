package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"pixelichi.com/util"
)

var testQueries *Queries
var store *Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Unable to load config. Error: ", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(conn)
	store = &Store{
		db:      conn,
		Queries: testQueries,
	}

	os.Exit(m.Run())
}
