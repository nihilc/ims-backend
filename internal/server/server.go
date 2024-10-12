package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nihilc/ims-backend/config"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	s := Server{
		port: config.Env.Port,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
