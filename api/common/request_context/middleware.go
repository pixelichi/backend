package request_context

import (
	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common"
)

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
