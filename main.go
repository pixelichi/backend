package main

import (
	_ "github.com/lib/pq"
	"shinypothos.com/api"
	"shinypothos.com/api/common"
	"shinypothos.com/api/data/ostore"
	"shinypothos.com/db"
	"shinypothos.com/token"
	"shinypothos.com/util"
)

// Define an alias for the Server type.
type Server = common.Server

func main() {
	config := util.LoadConfig(".")
	db := db.GetDBStore(config.DBDriver, config.DBSource)
	tokenMaker := token.NewPasetoMaker(config.TokenSymmetricKey)
	objectStore := ostore.NewMinioObjectStore(config.MINIO_ENDPOINT, config.MINIO_ACCESS_KEY, config.MINIO_SECRET_KEY)
	server := api.NewServer(config, db, &tokenMaker, &objectStore)
	server.Start(config.ServerAddress)
}
