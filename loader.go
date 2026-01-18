package postgres

import (
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	libsql "github.com/webcore-go/lib-sql"
	"github.com/webcore-go/webcore/app/config"
	"github.com/webcore-go/webcore/app/loader"
)

type PostgresLoader struct {
	name string
}

func (a *PostgresLoader) SetName(name string) {
	a.name = name
}

func (a *PostgresLoader) Name() string {
	return a.name
}

func (l *PostgresLoader) Init(args ...any) (loader.Library, error) {
	config := args[1].(config.DatabaseConfig)
	dsn := libsql.BuildDSN(config)

	db := &libsql.SQLDatabase{}

	driver := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	dialect := pgdialect.New()

	// Set up Bun SQL database wrapper
	db.SetBunDB(driver, dialect)

	err := db.Install(args...)
	if err != nil {
		return nil, err
	}

	db.Connect()

	return db, nil
}
