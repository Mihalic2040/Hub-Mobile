package deamon

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/Mihalic2040/Hub-Mobile/protocols"
	"google.golang.org/grpc"
)

func (s *Service) Start() {
	s.contex = context.Background()

	s.host = *CreateInstance(s.contex, s.cfg)
	s.dht = CreateDht(s.contex, s.host, s.cfg)
	Boot(s.contex, s.cfg, s.host)
	Bootstrap(s.contex, s.dht)
	//go DiscoverPeers(s.contex, s.host) //Global dht

	//log.Println(s.host.Peerstore().Peers().String())

	s.grcpServer = grpc.NewServer()

	protocols.RegisterPeersServer(s.grcpServer, s)
	protocols.RegisterRequestServiceServer(s.grcpServer, s)
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Println("failed to listen:", err)
	}
	go func() {
		if err := s.grcpServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

func (s *Service) Config(jsonStr string) {
	config := Config{}
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		log.Println("failed to deserialize JSON:", err)
	}

	s.cfg = config

	//log.Println("Config", s.cfg)
}

func (s *Service) RequestTest() {

	r := &protocols.RequestData{
		Peer:    "12D3KooWA4WULPc9yXQajeY8eoY7PgTgLKc72ENyNofswScdbQ7R",
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
	log.Println("Response:", response)
}
