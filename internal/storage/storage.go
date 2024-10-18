package storage

import (
	"fmt"
	"log"

	"github.com/nihilc/ims-backend/config"
	"github.com/nihilc/ims-backend/internal/storage/mysql"
	"github.com/nihilc/ims-backend/internal/storage/postgres"
)

var ErrUnsuportedDatabaseType = fmt.Errorf("unsupported database type")

const (
	Mysql    dbType = "mysql"
	Postgres dbType = "postgres"
)

type dbType string

type Storage struct {
}

func (s *Storage) Close() {
}

func NewStorage() (*Storage, error) {
	dbType := dbType(config.Env.DBType)
	switch dbType {
	case Mysql:
		_, err := mysql.NewMySQLConnection(nil)
		if err != nil {
			return nil, err
		}
		log.Printf("DB %s: Successfully connected!", dbType)
		return &Storage{}, nil
	case Postgres:
		_, err := postgres.NewPostgresConnection(nil)
		if err != nil {
			return nil, err
		}
		log.Printf("DB %s: Successfully connected!", dbType)
		return &Storage{}, nil
	default:
		return nil, ErrUnsuportedDatabaseType
	}
}
