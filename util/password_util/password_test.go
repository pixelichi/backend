package password_util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"shinypothos.com/util"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(20)

	hashedPass, err := hashPassword(password)
	require.NoError(t, err)
	require.NoError(t, checkPassword(password, hashedPass))

	hashedPass2, err := hashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPass, hashedPass2)
	require.NoError(t, checkPassword(password, hashedPass2))

	wrongPassword := util.RandomString(20)
	err = checkPassword(wrongPassword, hashedPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	err = checkPassword(wrongPassword, hashedPass2)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
