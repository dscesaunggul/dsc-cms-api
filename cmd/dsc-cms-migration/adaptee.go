package main

import (
	"context"
	"time"

	"github.com/dscueu/dsc-cms-api/pkg/migration"
	"github.com/jmoiron/sqlx"
)

// PGConnAdaptee is adaptee between dbconn migration and sqlx
type PGConnAdaptee struct {
	db *sqlx.DB
}

// Exec adaptee for db migration
func (p PGConnAdaptee) Exec(query string) error {
	_, err := p.db.Exec(query)
	return err
}

// Get adaptee for db migration
func (p PGConnAdaptee) Get(query string, data *migration.Data) error {
	res := p.db.QueryRowx(query)
	return res.StructScan(data)
}

// Trans adaptee for db migration
func (p PGConnAdaptee) Trans(cb func(txExec func(query string) error) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if nil != err {
		return err
	}
	err = cb(func(query string) error {
		_, err := tx.Exec(query)
		return err
	})
	if nil != err {
		return tx.Rollback()
	}

	return tx.Commit()
}

// NewPGConnAdaptee Initialize PGConnAdaptee
func NewPGConnAdaptee(db *sqlx.DB) migration.DBConn {
	return PGConnAdaptee{
		db: db,
	}
}
