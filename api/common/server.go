package common

import (
	"log"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/data/ostore"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/token"
	"shinypothos.com/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Config      util.Config
	DB          *db.Store
	Router      *gin.Engine
	TokenMaker  *token.Maker
	ObjectStore *ostore.OStore
}

// Start the http server on a specific address
func (server *Server) Start(address string) {
	err := server.Router.Run(address)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}
}
