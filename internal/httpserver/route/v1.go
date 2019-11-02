package route

import (
	"github.com/dscueu/dsc-cms-api/internal/controller"
	"github.com/dscueu/dsc-cms-api/internal/httpserver/container"
	"github.com/go-chi/chi"
)

// V1 return all route version one for dsc cms api
func V1(r chi.Router, c *container.ServiceContainer) {
	appinfo := controller.AppInfo{
		Service: c.AppInfo,
	}

	r.Get("/ping", appinfo.Ping)
}
