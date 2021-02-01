package rpc

import (
	"context"
	"log"

	pb "github.com/jmacelroy/redirect-demo/loot_pb"
)

type Version int

const (
	Version1 = iota
	Version2
)

func (v Version) String() string {
	return [...]string{"Version1", "Version2"}[v]
}

type Server struct {
	pb.UnimplementedLootRPCServer

	APIVersion Version
}

func (s *Server) LootDescription(ctx context.Context, req *pb.LootDescriptionRequest) (*pb.LootDescriptionResponse, error) {
	switch s.APIVersion {
	case Version1:
		log.Println("returning v1 loot data")
		return &pb.LootDescriptionResponse{
			Name:           "LEDx",
			Type:           "Medical Equipment",
			Weight:         0.23,
			GridSize:       "2x1",
			LootExperience: 50,
		}, nil
	default:
		log.Println("returning v2 loot data")
		return &pb.LootDescriptionResponse{
			Name:             "LEDx",
			Type:             "Medical Equipment",
			Weight:           0.23,
			GridSize:         "2x1",
			LootExperience:   50,
			Rarity:           "legendary",
			AveragePrice_24H: 710000.23,
			AveragePrice_7D:  805000.79,
		}, nil
	}
}
