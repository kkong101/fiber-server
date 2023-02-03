package repository

import (
	"context"
	"database/sql"
	sqlc "github.com/kkong101/fiber-server/db/sqlc"
)

type UserRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
	ctx     context.Context
}

func (u *UserRepository) createUser() error {

	tx, err := u.db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	qtx := u.queries.WithTx(tx)

	err = qtx.CreateUser(u.ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
