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

func GetUserFromIDOrAbort(c *gin.Context, db *db.Store, userId int64) (*db.User, error) {
	user, err := db.GetUserFromId(c, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, server_error.NewInvalidCredentialsError("Couldn't find user"))
			return nil, err
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
		return nil, err
	}

	return &user, nil
}

func GetUserOrAbort(c *gin.Context, db *db.Store, username string) (*db.User, error) {
	user, err := db.GetUserFromUsername(c, username)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, server_error.NewInvalidCredentialsError("Couldn't find user - "+username))
			return nil, err
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
		return nil, err
	}

	return &user, nil
}

func CreateUserOrAbort(c *gin.Context, db *db.Store, createUserParams *db.CreateUserParams) (*db.User, error) {
	user, err := db.CreateUser(c, *createUserParams)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			code_name := pqErr.Code.Name()
			msg := fmt.Sprintf("Error Code: %v, Error Text: %v", code_name, err.Error())

			switch code_name {

			case "unique_violation":
				err := server_error.NewForbiddenError("User already exists, please use a different username and email.")
				c.AbortWithStatusJSON(http.StatusForbidden, err)
				return nil, err

			default:
				err := server_error.NewInternalServerError(msg)
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return nil, err
			}
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
		return nil, err
	}

	return &user, nil
}
