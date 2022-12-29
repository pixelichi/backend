package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Begin Transaction
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	// Create a new query using sqlc lib, tx implemnents the DBTX interface
	q := New(tx)

	// Launch Query
	err = fn(q)

	// If the query fails
	if err != nil {

		// Attempt to roll back to prevent data corruption
		rbErr := tx.Rollback()

		// If we get an error when attempting to roll back
		if rbErr != nil {
			return fmt.Errorf("tx erro: %v, rollBack err: %v", err, rbErr)
		}

		// Query still failed so we need to return error
		return err
	}

	// If everything went well, we commit our transactions
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performance money transfer in one transaction
// - Create Transfer Record
// - Add Account Entry
// - Update Accounts Balance
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// Run everything under one transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		// Create From_Entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		// Create To_Entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// TODO - Need to make sure that transactions are always done in a repeatable a consistent order in order to minimize potential of deadlock

		// Update From Account and return new account object
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})

		if err != nil {
			return err
		}

		// Update To Account and return new account object
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}

		// everything went well, return err = nill
		return nil
	})

	return result, err
}
