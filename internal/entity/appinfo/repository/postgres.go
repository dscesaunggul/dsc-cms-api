package repository

import "github.com/jmoiron/sqlx"

// AppInfo centralize for command and query
type AppInfo struct {
	DBConn *sqlx.DB
}

// SelectOne ensuring the database can be query
func (a AppInfo) SelectOne() error {
	_, err := a.DBConn.Query("SELECT 1;")
	return err
}
