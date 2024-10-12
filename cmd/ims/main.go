package main

import (
	"log"

	"github.com/nihilc/ims-backend/internal/server"
)

func main() {
	s := server.NewServer()

	log.Printf("Listening on %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Error can't start server: %s", err)
	}
}
