package main

import (
	"log"

	"github.com/dscueu/dsc-cms-api/internal/httpserver"
	_ "github.com/lib/pq"
)

func main() {
	err := httpserver.Serve()
	if nil != err {
		log.Fatal(err.Error())
	}

	return
}
