package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	client "github.com/jmacelroy/redirect-demo/client-api"
)

func main() {
	lootDataEndpoint, ok := os.LookupEnv("LOOT_DATA_ENDPOINT")
	if !ok {
		log.Fatal("LOOT_DATA_ENDPOINT env var required")
	}
	server := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      client.DefaultMux(lootDataEndpoint),
	}

	fmt.Println("listening on :8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error setting up server: %+v\n", err)
	}
}
