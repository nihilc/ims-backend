package storage

import (
	"fmt"
	"log"

	"github.com/nihilc/ims-backend/config"
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

func NewStorage() (*Storage, error) {
	dbType := dbType(config.Env.DBType)
	switch dbType {
	case Mysql:
		return nil, nil
	case Postgres:
		_, err := postgres.NewPostgresStorage(postgres.PostgresConfig{
			Username: config.Env.DBUsername,
			Password: config.Env.DBPassword,
			Host:     config.Env.DBHost,
			Port:     config.Env.DBPort,
			Name:     config.Env.DBName,
		})
		if err != nil {
			return nil, err
		}
		log.Printf("DB %s: Successfully connected!", dbType)
		return &Storage{}, nil
	default:
		return nil, ErrUnsuportedDatabaseType
	}
}
