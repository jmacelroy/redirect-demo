package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lootdata "github.com/jmacelroy/redirect-demo/loot-data"
)

func main() {
	fullDataEnv := os.Getenv("DEMO_ENABLE_FULL_DATA")
	var enableFullData bool
	if fullDataEnv != "" {
		enableFullData = true
	}

	server := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8081",
		Handler:      lootdata.DefaultMux(enableFullData),
	}

	fmt.Println("loot data server listening on :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error setting up server: %+v\n", err)
	}
}
