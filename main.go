package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tech42-org/gh-merge-queue/internal/pets"
)

func main() {
	store := pets.NewStore()
	handler := pets.NewHandler(store)

	mux := http.NewServeMux()
	handler.Register(mux)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := ":8080"

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
