package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(20)

	hashedPass, err := HashPassword(password)
	require.NoError(t, err)
	require.NoError(t, CheckPassword(password, hashedPass))

	hashedPass2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPass, hashedPass2)
	require.NoError(t, CheckPassword(password, hashedPass2))

	wrongPassword := RandomString(20)
	err = CheckPassword(wrongPassword, hashedPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	err = CheckPassword(wrongPassword, hashedPass2)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
