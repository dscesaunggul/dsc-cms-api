package env

import (
	"os"
)

const (
	// HTTPListenPort specify port to be serve
	HTTPListenPort = "HTTP_LISTEN_PORT"
	// PGDBHost hostname postgres
	PGDBHost = "PG_DB_HOST"
	// PGDBPort port postgres
	PGDBPort = "PG_DB_PORT"
	// PGDBName dbname postgres
	PGDBName = "PG_DB_NAME"
	// PGDBUser user postgres
	PGDBUser = "PG_DB_USER"
	// PGDBPass user pass postgres
	PGDBPass = "PG_DB_PASS"
)

var envs = map[string]interface{}{
	HTTPListenPort: ":3000",
	PGDBHost:       "127.0.0.1",
	PGDBPort:       "5432",
	PGDBUser:       "postgres",
	PGDBPass:       "postgres",
	PGDBName:       "dscueucms",
}

// Get value from env, but if not found then take the value from default value at the variable above
func Get(key string) interface{} {
	v := os.Getenv(key)
	if "" == v {
		v, valid := envs[key]
		if valid {
			return v
		}
	}

	return v
}
