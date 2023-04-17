package account

// import (
// 	"database/sql"
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"shinypothos.com/token"
// )

// type getAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) getAccount(ctx *gin.Context) {
// 	var req getAccountRequest

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {

// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	if account.OwnerID != authPayload.UserID {
// 		err := errors.New("account does not belong to the authenticated user")
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }