package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"pixelichi.com/util"
)

func createRandomUser(t *testing.T) User {
	password := util.RandomString(10)
	hashedPass, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPass,
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)

	require.True(t, user.PasswordChangedAt.IsZero())

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Username, user.Username)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user2, err := testQueries.GetUserFromId(context.Background(), user.ID)
	require.NoError(t, err)

	user3, err := testQueries.GetUserFromUsername(context.Background(), user.Username)
	require.NoError(t, err)

	require.EqualValues(t, user2, user3)
	require.EqualValues(t, user, user3)
}
