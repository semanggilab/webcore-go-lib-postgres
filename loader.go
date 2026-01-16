package postgres

import (
	sql "github.com/webcore-go/lib-sql"
	"github.com/webcore-go/webcore/app/config"
	"github.com/webcore-go/webcore/app/loader"
	"gorm.io/driver/postgres"
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
	dsn := sql.BuildDSN(config)

	db := &sql.SQLDatabase{}
	db.SetDialect(postgres.Open(dsn))
	err := db.Install(args...)
	if err != nil {
		return nil, err
	}

	db.Connect()

	// l.DB = db
	return db, nil
}
