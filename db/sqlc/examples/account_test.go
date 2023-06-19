// package db

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// 	"shinypothos.com/util"
// )

// func createRandomAccount(t *testing.T) Account {
// 	user := createRandomUser(t)

// 	arg := CreateAccountParams{
// 		OwnerID:  user.ID, // randomly generated?
// 		Balance:  util.RandomBalance(),
// 		Currency: util.RandomCurrency(),
// 	}

// 	account, err := testQueries.CreateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account)

// 	require.Equal(t, arg.OwnerID, account.OwnerID)
// 	require.Equal(t, arg.Balance, account.Balance)
// 	require.Equal(t, arg.Currency, account.Currency)

// 	require.NotZero(t, account.ID)
// 	require.NotZero(t, account.CreatedAt)

// 	return account
// }

// func TestCreateAccount(t *testing.T) {
// 	createRandomAccount(t)
// }

// func TestGetAccount(t *testing.T) {
// 	acc := createRandomAccount(t)
// 	fetchedAcc, err := testQueries.GetAccount(context.Background(), acc.ID)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, fetchedAcc)

// 	require.Equal(t, acc.ID, fetchedAcc.ID)
// 	require.Equal(t, acc.Balance, fetchedAcc.Balance)
// 	require.Equal(t, acc.Currency, fetchedAcc.Currency)
// 	require.Equal(t, acc.OwnerID, fetchedAcc.OwnerID)
// 	require.WithinDuration(t, acc.CreatedAt, fetchedAcc.CreatedAt, time.Second)
// }

// func TestUpdateAccount(t *testing.T) {
// 	acc := createRandomAccount(t)

// 	newBalance := util.RandomBalance()

// 	args := UpdateAccountParams{
// 		ID:      acc.ID,
// 		Balance: newBalance,
// 	}

// 	newAcc, err := testQueries.UpdateAccount(context.Background(), args)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, acc)
// 	require.Equal(t, newAcc.Balance, newBalance)

// 	require.Equal(t, newAcc.ID, acc.ID)
// 	require.Equal(t, newAcc.Currency, acc.Currency)
// 	require.Equal(t, newAcc.OwnerID, acc.OwnerID)
// 	require.WithinDuration(t, newAcc.CreatedAt, acc.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	acc := createRandomAccount(t)

// 	require.NotEmpty(t, acc)

// 	testQueries.DeleteAccount(context.Background(), acc.ID)

// 	fetchedAcc, err := testQueries.GetAccount(context.Background(), acc.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, fetchedAcc)
// }

// func TestListAccounts(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomAccount(t)
// 	}

// 	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
// 		Limit:  5,
// 		Offset: 5,
// 	})

// 	require.NoError(t, err)
// 	require.Equal(t, len(accounts), 5)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 	}
// }
