package postgres

import (
	"fmt"

	"github.com/dscueu/dsc-cms-api/internal/config/env"
	"github.com/jmoiron/sqlx"
)

// Connect to postgres database
func Connect() (*sqlx.DB, error) {
	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.Get(env.PGDBHost), env.Get(env.PGDBPort), env.Get(env.PGDBUser), env.Get(env.PGDBPass), env.Get(env.PGDBName))
	db, err := sqlx.Open("postgres", ds)
	if nil != err {
		return nil, err
	}

	return db, nil
}
