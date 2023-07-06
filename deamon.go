package deamon

import (
	"context"
	"encoding/json"
	"log"
	"net"

	hub "github.com/Mihalic2040/Hub"
	"github.com/Mihalic2040/Hub-Mobile/protocols"
	"github.com/Mihalic2040/Hub/src/types"
	"google.golang.org/grpc"
)

type Server struct {
	protocols.UnimplementedPeersServer
}

type Service struct {
	hub.App
}

var app hub.App

func (s *Server) SayHello(ctx context.Context, req *protocols.PeersResponse) (*protocols.PeersResponse, error) {
	// Implement the logic to handle the gRPC request and return a response
	return &protocols.PeersResponse{Peers: app.Host.Peerstore().Peers().String()}, nil
}

func (s *Service) Start() {
	s.App.Start(false)

	// Create a gRPC server
	server := grpc.NewServer()

	myServer := &Server{}

	// Register your gRPC service implementation
	protocols.RegisterPeersServer(server, myServer)

	// Create a listener for your chosen network protocol
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	app = s.App
	// Start the gRPC server
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

func (s *Service) Config(jsonStr string) {
	// Create a new instance of the Config struct
	config := types.Config{}

	// Unmarshal the JSON string into the Config struct
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		log.Println("failed to deserialize JSON:", err)
	}

	// Set the deserialized Config in the App struct
	s.App.Settings(config)

}
