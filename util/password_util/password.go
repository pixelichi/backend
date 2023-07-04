package password_util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"shinypothos.com/api/common/server_error"
)

// returns the bcrypt hash of the password
func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPass), nil
}

func HashPasswordOrAbort(c *gin.Context, password string) string {
	pass, err := hashPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
	}
	return pass
}

// Checks if a provided password is correct
func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckPasswordOrAbort(c *gin.Context, pass string, hashedPass string) (error) {
	err := checkPassword(pass, hashedPass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, server_error.NewInvalidCredentialsError("Password provided is incorrect. Try again."))
		return err
	}

	return nil
}
