package request_context

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common"
	"shinypothos.com/api/common/server_error"
	"shinypothos.com/api/data/ostore"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/token"
	"shinypothos.com/util"
)

const requestContextKey = "request_context"

type RequestContext struct {
	Config     *util.Config
	DB         *db.Store
	TokenMaker *token.Maker
	OS         *ostore.OStore
}

func SetReqCtx(s *common.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqContext := RequestContext{
			Config:     &s.Config,
			DB:         s.DB,
			OS:         s.ObjectStore,
			TokenMaker: s.TokenMaker,
		}

		c.Set(requestContextKey, &reqContext)
		c.Next()
	}
}

func getReqCtx(c *gin.Context) (*RequestContext, error) {
	reqContext, exists := c.Get(requestContextKey)
	if !exists {
		return &RequestContext{}, errors.New("Could not find " + requestContextKey + " within context")
	}

	// Convert "any" to RequestContext
	rc := reqContext.(*RequestContext)

	return rc, nil
}

func GetReqCtxOrInternalServerError(c *gin.Context) *RequestContext {
	reqCtx, err := getReqCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Unable to process request context"))
	}

	return reqCtx
}
