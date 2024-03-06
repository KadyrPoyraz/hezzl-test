package db

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresqlDB struct {
	*sqlx.DB
}

type PostgresqlTx struct {
	Txm *sqlx.Tx
	context.Context
}

func NewPostgresqlDB(dsn string) (*PostgresqlDB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	postgresqlDb := &PostgresqlDB{db}

	return postgresqlDb, nil
}

func (db *PostgresqlDB) StartTransaction(ctx context.Context) (Transaction, error) {
	tx, err := db.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &PostgresqlTx{
		Txm:     tx,
		Context: ctx,
	}, nil
}

func (tx *PostgresqlTx) Tx() *sqlx.Tx {
	return tx.Txm
}

func (tx *PostgresqlTx) Commit() error {
    return tx.Commit()
}

func (tx *PostgresqlTx) Rollback() error {
    return tx.Rollback()
}
