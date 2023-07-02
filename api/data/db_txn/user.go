package db_txn

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"shinypothos.com/api/common/server_error"
	db "shinypothos.com/db/sqlc"
)

func GetUserOrAbort(c *gin.Context, db *db.Store, username string) db.User {
	user, err := db.GetUserFromUsername(c, username)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, server_error.NewInvalidCredentialsError("Couldn't find user - "+username))
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
	}

	return user
}

func CreateUserOrAbort(c *gin.Context, db *db.Store, createUserParams *db.CreateUserParams) db.User {
	user, err := db.CreateUser(c, *createUserParams)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			code_name := pqErr.Code.Name()
			msg := fmt.Sprintf("Error Code: %v, Error Text: %v", code_name, err.Error())

			switch code_name {

			case "unique_violation":
				c.AbortWithStatusJSON(http.StatusForbidden, server_error.NewForbiddenError("User already exists, please use a different username and email."))
				break

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(msg))
				break
			}
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
	}

	return user
}
