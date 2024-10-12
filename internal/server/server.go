package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nihilc/ims-backend/config"
	"github.com/nihilc/ims-backend/internal/storage"
)

type Server struct {
	port int
	db   *storage.Storage
}

func NewServer(db *storage.Storage) *http.Server {
	s := Server{
		port: config.Env.Port,
		db:   db,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
