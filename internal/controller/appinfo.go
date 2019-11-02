package controller

import (
	"net/http"

	"github.com/dscueu/dsc-cms-api/internal/entity/appinfo/service"
	"github.com/dscueu/dsc-cms-api/pkg/jsonhttp"
)

// AppInfo get information about app like version and health check
type AppInfo struct {
	Service *service.AppInfo
}

// Ping ensure the application can be run and execute the database query
func (a AppInfo) Ping(w http.ResponseWriter, r *http.Request) {
	status := a.Service.Ping()
	jsonhttp.Encode(w, http.StatusOK, map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, map[string]bool{
		"status": status,
	})
}
