package migration

import "fmt"

// Data is all column been specify in migration table
type Data struct {
	Value int `db:"value"`
}

// DBConn is interface for db connection
type DBConn interface {
	// Exec for DDL & DML (Select Invalid)
	Exec(query string) error
	// Get for DML (Select)
	Get(query string, data *Data) error
	// Trans for DDL Or DML transaction
	Trans(func(txExec func(query string) error) error) error
}

// Migration need the connection db to execute it
type Migration struct {
	DBConn DBConn
}

// Run schemas given
func (m Migration) Run(schemas []string) error {
	err := m.createTblMigrationIfNoExists()
	if nil != err {
		return err
	}

	// if err is not empty we assume the data is null at the table migration version, so we just running the
	// schema and add version
	curr := len(schemas)
	old, err := m.getOldVersion()
	if nil != err {
		err = m.do(schemas)
		if nil != err {
			return err
		}

		return m.addVersion(curr)
	}

	if curr > old {
		err = m.do(schemas)
		if nil != err {
			return err
		}

		return m.updateVersion(old, curr)
	}

	return nil
}

func (m Migration) createTblMigrationIfNoExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS migration_version (
		value smallint NOT NULL
	);
	`
	return m.DBConn.Exec(query)
}

func (m Migration) getOldVersion() (int, error) {
	d := new(Data)
	query := `SELECT * FROM migration_version LIMIT 1;`
	err := m.DBConn.Get(query, d)

	return d.Value, err
}

func (m Migration) addVersion(version int) error {
	query := fmt.Sprintf(`INSERT INTO migration_version (value) VALUES (%d);`, version)
	return m.DBConn.Exec(query)
}

func (m Migration) updateVersion(k int, version int) error {
	query := fmt.Sprintf(`UPDATE migration_version SET value='%d' WHERE value='%d';`, k, version)
	return m.DBConn.Exec(query)
}

func (m Migration) do(schemas []string) error {
	return m.DBConn.Trans(func(txExec func(query string) error) error {
		for _, schema := range schemas {
			err := txExec(schema)
			if nil != err {
				return err
			}
		}

		return nil
	})
}
