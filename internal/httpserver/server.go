package httpserver

import (
	"fmt"
	"net/http"

	"github.com/dscueu/dsc-cms-api/internal/config/env"
	"github.com/dscueu/dsc-cms-api/internal/httpserver/container"
	"github.com/dscueu/dsc-cms-api/internal/httpserver/route"
	"github.com/go-chi/chi"
)

// Serve HTTP Server for DSC Content Management System
func Serve() error {
	r := chi.NewRouter()
	c := container.CreateServiceContainer()

	// specify the route for specific version at the below
	r.Route("/v1", func(r chi.Router) {
		route.V1(r, c)
	})

	port := env.Get(env.HTTPListenPort).(string)
	fmt.Printf("Starting server at %s\n", port)
	return http.ListenAndServe(port, r)
}
