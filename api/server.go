package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pixelichi.com/api/common"
	"pixelichi.com/api/user"
	db "pixelichi.com/db/sqlc"
	"pixelichi.com/token"
	"pixelichi.com/util"
)

// Define an alias for the Server type.
type Server = common.Server

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}

	setupRouter(config, server)

	return server, nil
}

func setupRouter(config util.Config, server *Server) {
	router := gin.Default()

  // CORS for https://foo.com and https://github.com origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // - Credentials share
  // - Preflight requests cached for 12 hours
  router.Use(cors.New(cors.Config{
    AllowMethods:     []string{"POST", "GET"},
    AllowHeaders:     []string{"Origin","Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, config.ALLOW_ORIGIN)
    },
    MaxAge: 12 * time.Hour,
  }))

	// No auth needed
	// router.POST("/users", server.createUser)
	router.POST("/users/login", withServerContext(server, user.LoginUser))

	// add routes to router
	// authRoutes := router.Group("/").Use(authMiddleware(server.TokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	server.Router = router
}

// Allows us to pass the server to the handler functions
func withServerContext(server *Server, handler func(server *Server, c *gin.Context)) func(c *gin.Context) {
	return func (c *gin.Context) {
		handler(server, c)
	}
}
