package storage

import (
	"fmt"

	m "github.com/golang-migrate/migrate/v4"
	mDatabase "github.com/golang-migrate/migrate/v4/database"
	mMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	mPostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/nihilc/ims-backend/config"
	"github.com/nihilc/ims-backend/internal/storage/mysql"
	"github.com/nihilc/ims-backend/internal/storage/postgres"
)

func NewMigrate() (*m.Migrate, error) {
	var (
		dbType       = dbType(config.Env.DBType)
		sourceURL    = fmt.Sprintf("file://internal/storage/%s/migrations", dbType)
		databaseName = string(dbType)
		driver       mDatabase.Driver
	)

	switch dbType {
	case Mysql:
		db, err := mysql.NewMySQLConnection(nil)
		if err != nil {
			return nil, err
		}
		driver, err = mMysql.WithInstance(db, &mMysql.Config{})
		if err != nil {
			return nil, err
		}
	case Postgres:
		db, err := postgres.NewPostgresConnection(nil)
		if err != nil {
			return nil, err
		}
		driver, err = mPostgres.WithInstance(db, &mPostgres.Config{})
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnsuportedDatabaseType
	}

	return m.NewWithDatabaseInstance(
		sourceURL,
		databaseName,
		driver,
	)
}
