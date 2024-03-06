package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Commit() error
	Rollback() error
    Tx() *sqlx.Tx
}

type DB interface {
	StartTransaction(context.Context) (Transaction, error)
}
