package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tech42-org/gh-merge-queue/internal/pets"
)

func main() {
	store := pets.NewStore()
	handler := pets.NewHandler(store)

	mux := http.NewServeMux()
	handler.Register(mux)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("pets api listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
