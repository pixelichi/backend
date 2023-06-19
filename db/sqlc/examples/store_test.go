// package db

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"shinypothos.com/util"
// )

// func TestTransferTx(t *testing.T) {
// 	fromAcc := createRandomAccount(t)
// 	toAcc := createRandomAccount(t)
// 	transferAmount := util.RandomBalance()

// 	transferResult, err := store.TransferTx(context.Background(), TransferTxParams{
// 		FromAccountID: fromAcc.ID,
// 		ToAccountID:   toAcc.ID,
// 		Amount:        transferAmount,
// 	})

// 	require.NoError(t, err)
// 	require.NotEmpty(t, transferResult)

// 	fromAccPostTransfer, err := testQueries.GetAccount(context.Background(), fromAcc.ID)
// 	require.NoError(t, err)

// 	toAccPostTransfer, err := testQueries.GetAccount(context.Background(), toAcc.ID)
// 	require.NoError(t, err)

// 	require.EqualValues(t, fromAccPostTransfer, transferResult.FromAccount)
// 	require.EqualValues(t, toAccPostTransfer, transferResult.ToAccount)

// 	// Check that the new balance on the fromAccount is deducted the transfer balance
// 	require.Equal(t, transferResult.FromAccount.Balance, fromAcc.Balance-transferAmount)

// 	// Check that the new balance on the toAccount is increased by the transfer balance
// 	require.Equal(t, transferResult.ToAccount.Balance, toAcc.Balance+transferAmount)

// 	// Check that the FromEntry has the right data
// 	require.Equal(t, transferResult.FromEntry.AccountID, fromAcc.ID)
// 	require.Equal(t, transferResult.FromEntry.Amount, -transferAmount)

// 	// Check that the ToEntry has the right data
// 	require.Equal(t, transferResult.ToEntry.AccountID, toAcc.ID)
// 	require.Equal(t, transferResult.ToEntry.Amount, transferAmount)
// }

// func TestTransfersConcurrent(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)

// 	// Amount of concurrent threads
// 	n := 50

// 	// Amount to transfer per transaction
// 	amount := util.RandomBalance()

// 	errsCh := make(chan error)
// 	resultsCh := make(chan TransferTxResult)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			result, err := store.TransferTx(context.Background(), TransferTxParams{
// 				FromAccountID: account1.ID,
// 				ToAccountID:   account2.ID,
// 				Amount:        amount,
// 			})
// 			errsCh <- err
// 			resultsCh <- result
// 		}()
// 	}

// 	existed := make(map[int]bool)
// 	for i := 0; i < n; i++ {
// 		err := <-errsCh
// 		require.NoError(t, err)

// 		result := <-resultsCh

// 		require.NotEmpty(t, result)

// 		require.Equal(t, result.FromAccount.ID, account1.ID)
// 		require.Equal(t, result.ToAccount.ID, account2.ID)

// 		require.Equal(t, result.FromEntry.Amount, -amount)
// 		require.Equal(t, result.ToEntry.Amount, amount)

// 		require.Equal(t, result.Transfer.Amount, amount)
// 		require.NotZero(t, result.Transfer.ID)        // Since auto increment field, should not be 0
// 		require.NotZero(t, result.Transfer.CreatedAt) // Not zero since defaulted

// 		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)
// 		require.NoError(t, err)

// 		// Check From Entry
// 		var fromEntry Entry
// 		fromEntry, err = store.GetEntry(context.Background(), result.FromEntry.ID)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, fromEntry)
// 		require.Equal(t, result.FromEntry.Amount, -amount)
// 		require.NotZero(t, result.FromEntry.CreatedAt)
// 		require.EqualValues(t, fromEntry, result.FromEntry)

// 		// Check To Entry
// 		var toEntry Entry
// 		toEntry, err = store.GetEntry(context.Background(), result.ToEntry.ID)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, toEntry)
// 		require.Equal(t, result.ToEntry.Amount, amount)
// 		require.NotZero(t, result.ToEntry.CreatedAt)
// 		require.EqualValues(t, toEntry, result.ToEntry)

// 		// Check Accounts
// 		fromAccount := result.FromAccount
// 		require.NotEmpty(t, fromAccount)
// 		require.Equal(t, fromAccount.ID, account1.ID)

// 		toAccount := result.ToAccount
// 		require.NotEmpty(t, toAccount)
// 		require.Equal(t, toAccount.ID, account2.ID)

// 		// Still need to verify balances
// 		diff1 := account1.Balance - fromAccount.Balance
// 		diff2 := toAccount.Balance - account2.Balance

// 		require.Equal(t, diff1, diff2)
// 		require.True(t, diff1 > 0)
// 		require.True(t, diff1%amount == 0) // The difference should be perfectly divisible by amount

// 		k := int(diff1 / amount)
// 		require.True(t, k >= 1 && k <= n)
// 		require.NotContains(t, existed, k)
// 		existed[k] = true
// 	}

// 	// Check the final updated balances
// 	totalAmount := int64(n) * amount

// 	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
// 	require.NoError(t, err)

// 	require.Equal(t, updatedAccount1.Balance, account1.Balance-totalAmount)
// 	require.Equal(t, updatedAccount2.Balance, account2.Balance+totalAmount)

// }
