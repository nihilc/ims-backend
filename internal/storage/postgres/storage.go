package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nihilc/ims-backend/config"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (cfg PostgresConfig) formatDNS() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
}

func NewPostgresStorage(cfg *PostgresConfig) (*sql.DB, error) {
	var db *sql.DB
	var err error
	// load default if there isn't config
	if cfg == nil {
		cfg = &PostgresConfig{
			Username: config.Env.DBUsername,
			Password: config.Env.DBPassword,
			Host:     config.Env.DBHost,
			Port:     config.Env.DBPort,
			Name:     config.Env.DBName,
		}
	}
	db, err = sql.Open("postgres", cfg.formatDNS())
	if err != nil {
		log.Printf("Error opening database: %s", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %s", err)
		return nil, err
	}
	return db, nil
}
