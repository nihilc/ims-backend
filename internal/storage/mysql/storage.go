package mysql

import (
	"database/sql"
	"fmt"
	"log"

	dmysql "github.com/go-sql-driver/mysql"
	"github.com/nihilc/ims-backend/config"
)

func NewMySQLConnection(cfg *dmysql.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error
	// load default if there isn't config
	if cfg == nil {
		cfg = &dmysql.Config{
			User:   config.Env.DBUsername,
			Passwd: config.Env.DBPassword,
			Addr:   fmt.Sprintf("%s:%s", config.Env.DBHost, config.Env.DBPort),
			DBName: config.Env.DBName,
		}
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
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
