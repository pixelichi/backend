package api

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shinypothos.com/api/auth"
	"shinypothos.com/api/common"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/data/ostore"
	"shinypothos.com/api/middleware"
	"shinypothos.com/api/user"
	db "shinypothos.com/db/sqlc"
	token "shinypothos.com/token"

	"shinypothos.com/util"
)

// Define an alias for the Server type.
type Server = common.Server

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, DB *db.Store, tokenMaker *token.Maker, objectStore *ostore.OStore) *Server {

	server := &Server{
		Config:      config,
		DB:          DB,
		TokenMaker:  tokenMaker,
		ObjectStore: objectStore,
	}

	setupRouter(config, server)
	return server
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
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, config.ALLOW_ORIGIN) || (util.IsLocalEnv(config) && strings.HasPrefix(origin, config.ALLOW_ORIGIN_LAN))
		},
		MaxAge: 12 * time.Hour,
	}))

	const userRoute = "user"

	// No auth needed
	noAuthRoutes := router.Group("/").Use(request_context.SetReqCtx(server))

	noAuthRoutes.GET("/auth/check", auth.CheckAuth)
	noAuthRoutes.POST("/"+userRoute+"/login", user.LoginUser)

	if util.IsLocalEnv(config) {
		noAuthRoutes.POST("/"+userRoute+"/sign_up", user.SignUp)
	}

	// add routes to router
	authRoutes := router.Group("/").
	Use(middleware.AuthMiddleware(*server.TokenMaker)).
	Use(request_context.SetReqCtx(server))

	authRoutes.POST("/"+userRoute+"/set_profile_photo", user.SetProfilePicture)
	authRoutes.GET("/"+userRoute+"/get_profile_photo", user.GetProfilePicture)

	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.POST("/accounts", server.createAccount)

	server.Router = router
}

// Allows us to pass the server to the handler functions
// func withServerContext(server *Server, handler func(server *Server, c *gin.Context)) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		handler(server, c)
// 	}
// }
