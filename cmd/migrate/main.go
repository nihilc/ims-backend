package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nihilc/ims-backend/config"
	"github.com/nihilc/ims-backend/internal/storage"
)

func main() {
	if len(os.Args) < 2 {
		handleHelp()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "up":
		handleUp()
	case "down":
		handleDown()
	case "new":
		handleNew(os.Args[2:])
	default:
		handleHelp()
	}
}

func handleHelp() {
	fmt.Println("Usage: migrate COMMAND [arg...]")
	fmt.Println("Commands:")
	fmt.Println("\tup")
	fmt.Println("\t\tApply all up migrations")
	fmt.Println("\tdown")
	fmt.Println("\t\tApply all down migrations")
	fmt.Println("\tnew [-db DB] NAME")
	fmt.Println("\t\tnew up/down migrations files titled NAME for $DB_TYPE database")
	fmt.Println("\t\tCan use -db option to specify the database for the new migration files")
}

func handleUp() {
	m, err := storage.NewMigrate()
	defer m.Close()
	if err != nil {
		log.Fatalf("Error can't create storage migrate: %s", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error can't apply up migrate: %s", err)
	}
}

func handleDown() {
	m, err := storage.NewMigrate()
	defer m.Close()
	if err != nil {
		log.Fatalf("Error can't create storage migrate: %s", err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error can't apply down migrate: %s", err)
	}
}

func handleNew(args []string) {
	var name, db string
	cmd := flag.NewFlagSet("new", flag.ExitOnError)
	cmd.StringVar(&db, "db", config.Env.DBType, "database for the new migration files")

	if err := cmd.Parse(args); err != nil {
		log.Fatalf("Error parsing args: %s", err)
	}
	if cmd.NArg() < 1 {
		log.Fatalf("Error please specify NAME for the new migration files")
	}
	name = cmd.Arg(0)

	if err := createMigrationFiles(name, db); err != nil {
		log.Fatalf("Error creating new migration files: %s", err)
	}
}

func createMigrationFiles(name, db string) error {
	timestamp := time.Now().Format("20060102150405")
	dir := fmt.Sprintf("internal/storage/%s/migrations", db)

	upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
	downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, name)

	upFilePath := filepath.Join(dir, upFileName)
	downFilePath := filepath.Join(dir, downFileName)

	upFile, err := os.Create(upFilePath)
	if err != nil {
		return fmt.Errorf("can't create up file: %s", err)
	}
	defer upFile.Close()

	downFile, err := os.Create(downFilePath)
	if err != nil {
		return fmt.Errorf("can't create down file: %s", err)
	}
	defer downFile.Close()

	fmt.Printf("Migration files created: \n%s\n%s\n", upFilePath, downFilePath)
	return nil
}
