package storage

import (
	"fmt"

	"github.com/nihilc/ims-backend/config"
)

var ErrUnsuportedDatabaseType = fmt.Errorf("unsupported database type")

const (
	Mysql    dbType = "mysql"
	Postgres dbType = "postgres"
)

type dbType string

type Storage struct {
}

func NewStorage() (*Storage, error) {
	dbType := dbType(config.Env.DBType)
	switch dbType {
	case Mysql:
		return nil, nil
	case Postgres:
		return nil, nil
	default:
		return nil, ErrUnsuportedDatabaseType
	}
}
