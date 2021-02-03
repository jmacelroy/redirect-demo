package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	client "github.com/jmacelroy/redirect-demo/client-api"
)

func main() {
	server := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      client.DefaultMux("loot-data.jacob:8081"),
	}

	fmt.Println("listening on :8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error setting up server: %+v\n", err)
	}
}
