package main

import (
	"log"

	"github.com/dscueu/dsc-cms-api/cmd/dsc-cms-migration/schema"
	"github.com/dscueu/dsc-cms-api/internal/config/postgres"
	"github.com/dscueu/dsc-cms-api/pkg/migration"
	_ "github.com/lib/pq"
)

var schemas = []string{
	schema.First,
}

func main() {
	pgconn, err := postgres.Connect()
	if nil != err {
		log.Fatal(err.Error())
	}

	dbconn := NewPGConnAdaptee(pgconn)
	m := migration.Migration{DBConn: dbconn}
	err = m.Run(schemas)
	if nil != err {
		log.Fatal(err.Error())
	}

	log.Println("Migration success")
}
