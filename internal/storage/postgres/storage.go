package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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

func NewPostgresStorage(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.formatDNS())
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
