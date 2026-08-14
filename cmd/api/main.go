package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cfotel2604/monitoring-api/internal/config"
	"github.com/cfotel2604/monitoring-api/internal/handler"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	handler.New().Register(mux)

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("monitoring-api listening on %s", cfg.HTTPAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
