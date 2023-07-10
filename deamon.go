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

type Service struct {
	App hub.App

	grcpServer *grpc.Server
	// server      *Server

	//grpc stuff
	protocols.UnimplementedPeersServer
	protocols.UnimplementedRequestServiceServer
}

func (s *Service) Start() {
	s.App.App = *s.App.Start(false)

	// s.initialize()
	// Create a gRPC server
	s.grcpServer = grpc.NewServer()

	// Register your gRPC service implementation
	protocols.RegisterPeersServer(s.grcpServer, s)
	protocols.RegisterRequestServiceServer(s.grcpServer, s)

	// Create a listener for your chosen network protocol
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	go func() {
		if err := s.grcpServer.Serve(lis); err != nil {
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

func (s *Service) RequestTest() {

	r := &protocols.RequestData{
		Peer:    "12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
		Handler: "MyHandler",
		Payload: "request",
	}
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost"+":"+"4545", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := protocols.NewRequestServiceClient(conn)

	// Prepare the request

	// Send the gRPC request
	response, err := client.Request(context.Background(), r)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	// Process the response
	log.Println(response)
}
