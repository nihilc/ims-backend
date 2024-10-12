package main

import (
	"log"

	"github.com/nihilc/ims-backend/internal/server"
	"github.com/nihilc/ims-backend/internal/storage"
)

func main() {
	db, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("Error can't connect database: %s", err)
	}

	s := server.NewServer(db)

	log.Printf("Listening on %s", s.Addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Error can't start server: %s", err)
	}
}
