package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/jmacelroy/redirect-demo/loot_pb"
	"google.golang.org/grpc"
)

func DefaultMux(rpcEndpoint string) *http.ServeMux {
	handlers := Handlers{rpcEndpoint: rpcEndpoint}
	mux := http.NewServeMux()
	mux.Handle("/loot/description", http.HandlerFunc(handlers.LootData))
	return mux
}

type Handlers struct {
	rpcEndpoint string
}

func (h Handlers) LootData(w http.ResponseWriter, r *http.Request) {
	names, ok := r.URL.Query()["name"]
	log.Printf("querying for description of %s\n", names)

	if !ok || len(names[0]) < 1 {
		http.Error(w, "name query parameter missing", http.StatusBadRequest)
		return
	}

	log.Printf("RPC Endpoint: %s\n", h.rpcEndpoint)
	conn, err := grpc.Dial(h.rpcEndpoint, grpc.WithInsecure())
	if err != nil {
		http.Error(w, fmt.Sprintf("rpc: %+v", err), http.StatusInternalServerError)
		return
	}
	client := pb.NewLootRPCClient(conn)
	resp, err := client.LootDescription(context.Background(), &pb.LootDescriptionRequest{Name: names[0]})
	if err != nil {
		http.Error(w, fmt.Sprintf("rpc: %+v", err), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}
