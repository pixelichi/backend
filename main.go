package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"shinypothos.com/api"
	"shinypothos.com/api/common"
	"shinypothos.com/api/minio"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/util"
)

// Define an alias for the Server type.
type Server = common.Server

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

	minio.InitMinioClient(config.MINIO_ENDPOINT, config.MINIO_ACCESS_KEY, config.MINIO_SECRET_KEY)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}
}
