package container

import (
	"log"

	"github.com/dscueu/dsc-cms-api/internal/config/postgres"
	"github.com/dscueu/dsc-cms-api/internal/entity/appinfo/service"
)

// ServiceContainer all services
type ServiceContainer struct {
	AppInfo *service.AppInfo
}

// CreateServiceContainer create and cache the container services
func CreateServiceContainer() *ServiceContainer {
	pgconn, err := postgres.Connect()
	if nil != err {
		log.Fatal(err)
	}

	return &ServiceContainer{
		AppInfo: service.NewAppInfo(pgconn),
	}
}
