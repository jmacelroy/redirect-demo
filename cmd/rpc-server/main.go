package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jmacelroy/redirect-demo/loot_pb"
	rpc "github.com/jmacelroy/redirect-demo/rpc-server"
	"google.golang.org/grpc"
)

func main() {
	rpcAPIVersion := os.Getenv("DEMO_RPC_API_VERSION")
	var rpcAPI rpc.Server
	if rpcAPIVersion == "" || rpcAPIVersion == "v1" {
		rpcAPI = rpc.Server{APIVersion: rpc.Version1}
	} else {
		rpcAPI = rpc.Server{APIVersion: rpc.Version2}
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLootRPCServer(grpcServer, &rpcAPI)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("rpc server listening on :8081...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
