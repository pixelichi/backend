package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"pixelichi.com/api"
	db "pixelichi.com/db/sqlc"
	"pixelichi.com/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://admin:password@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:1337"
// )

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}
}
