package storage

import (
	"fmt"
	"log"

	dmysql "github.com/go-sql-driver/mysql"
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

func NewStorage() (*Storage, error) {
	dbType := dbType(config.Env.DBType)
	switch dbType {
	case Mysql:
		_, err := mysql.NewMySQLStorage(dmysql.Config{
			User:   config.Env.DBUsername,
			Passwd: config.Env.DBPassword,
			Addr:   fmt.Sprintf("%s:%s", config.Env.DBHost, config.Env.DBPort),
			DBName: config.Env.DBName,
		})
		if err != nil {
			return nil, err
		}
		log.Printf("DB %s: Successfully connected!", dbType)
		return &Storage{}, nil
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
