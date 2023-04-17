package account

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lib/pq"
// 	"shinypothos.com/api/common"
// 	db "shinypothos.com/db/sqlc"
// 	"shinypothos.com/token"
// )

// type createAccountRequest struct {
// 	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
// }

// func (server *Server) createAccount(ctx *gin.Context) {
// 	var req createAccountRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(common.AuthorizationPayloadKey).(*token.Payload)

// 	arg := db.CreateAccountParams{
// 		OwnerID:  authPayload.UserID,
// 		Currency: req.Currency,
// 		Balance:  int64(0),
// 	}

// 	account, err := server.store.CreateAccount(ctx, arg)
// 	if err != nil {

// 		// If it's a DB error, let's choose how to respond
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}

// 		// Non DB error, still need to fail
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }