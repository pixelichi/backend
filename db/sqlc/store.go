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
