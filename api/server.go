package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "pixelichi.com/db/sqlc"
	"pixelichi.com/token"
	"pixelichi.com/util"
)

// Server serves http
type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

  // CORS for https://foo.com and https://github.com origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // - Credentials share
  // - Preflight requests cached for 12 hours
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"POST", "GET"},
    AllowHeaders:     []string{"Origin","Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "http://localhost:3000"
    },
    MaxAge: 12 * time.Hour,
  }))

	// No auth needed
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// add routes to router
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)


	server.router = router
}

// Start the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}