package mysql

import (
	"database/sql"
	"log"

	dmysql "github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg dmysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
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
