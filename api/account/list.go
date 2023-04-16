package account

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"pixelichi.com/api/common"
	"pixelichi.com/api/common/error"
	db "pixelichi.com/db/sqlc"
	"pixelichi.com/token"
)

type Server = common.Server


type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	OwnerID  int64 `json:"owner_id"`
}

func ListAccounts(server *Server, ctx *gin.Context) {
	var req listAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error.NewBadRequestError(err.Error()))
		return
	}

	authPayload := ctx.MustGet(common.AuthorizationPayloadKey).(*token.Payload)

	accounts, err := server.Store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OwnerID: authPayload.UserID,
	})
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, error.NewNotFoundError(err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
