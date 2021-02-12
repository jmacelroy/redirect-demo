package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	client "github.com/jmacelroy/redirect-demo/client-api"
	"github.com/justinas/alice"
)

func main() {
	lootDataEndpoint, ok := os.LookupEnv("LOOT_DATA_ENDPOINT")
	if !ok {
		log.Fatal("LOOT_DATA_ENDPOINT env var required")
	}

	handlers := client.Handlers{lootDataEndpoint}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.LootData))

	chain := alice.New(
		client.InjectDivertHeaderContext(),
	).Then(mux)

	server := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      chain,
	}

	fmt.Println("listening on :8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error setting up server: %+v\n", err)
	}
}
