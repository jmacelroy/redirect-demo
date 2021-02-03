package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	lootdata "github.com/jmacelroy/redirect-demo/loot-data"
)

func main() {
	server := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8081",
		Handler:      lootdata.DefaultMux(false),
	}

	fmt.Println("loot data server listening on :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error setting up server: %+v\n", err)
	}
}
