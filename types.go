package deamon

import (
	"context"

	"github.com/Mihalic2040/Hub-Mobile/protocols"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"google.golang.org/grpc"
)

type Config struct {
	Host       string
	Port       string
	Secret     string
	Rendezvous string
	ProtocolId string
	Bootstrap  string
	DHTServer  bool
}

type Service struct {
	host   host.Host
	dht    *dht.IpfsDHT
	cfg    Config
	contex context.Context

	grcpServer *grpc.Server
	// server      *Server

	//grpc stuff
	protocols.UnimplementedPeersServer
	protocols.UnimplementedRequestServiceServer
}
