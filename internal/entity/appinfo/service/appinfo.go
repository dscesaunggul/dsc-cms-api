package service

import (
	"github.com/dscueu/dsc-cms-api/internal/entity/appinfo/repository"
	"github.com/jmoiron/sqlx"
)

// AppInfo get information about app like version and health check
type AppInfo struct {
	repo *repository.AppInfo
}

// Ping check connection db
func (a AppInfo) Ping() bool {
	err := a.repo.SelectOne()
	if nil != err {
		return false
	}

	return true
}

// NewAppInfo initialize appinfo
func NewAppInfo(pgconn *sqlx.DB) *AppInfo {
	return &AppInfo{
		repo: &repository.AppInfo{
			DBConn: pgconn,
		},
	}
}
