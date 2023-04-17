package common

import (
	"github.com/gin-gonic/gin"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/token"
	"shinypothos.com/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Config     util.Config
	Store      *db.Store
	Router     *gin.Engine
	TokenMaker token.Maker
}

// Start the http server on a specific address
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}