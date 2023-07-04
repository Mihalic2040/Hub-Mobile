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
	protocols.UnimplementedMyServiceServer
}

type Service struct {
	hub.App
}

func (s *Server) SayHello(ctx context.Context, req *protocols.HelloRequest) (*protocols.HelloResponse, error) {
	// Implement the logic to handle the gRPC request and return a response
	return &protocols.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func (s *Service) Start() {
	s.App.Start(false)

	// Create a gRPC server
	server := grpc.NewServer()

	myServer := &Server{}

	// Register your gRPC service implementation
	protocols.RegisterMyServiceServer(server, myServer)

	// Create a listener for your chosen network protocol
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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

func (s *Service) Test() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:4545", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := protocols.NewMyServiceClient(conn)

	// Prepare the request
	req := &protocols.HelloRequest{
		Name: "John",
	}

	// Send the gRPC request
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}

	// Process the response
	log.Printf("Response: %s", resp.Message)
}
