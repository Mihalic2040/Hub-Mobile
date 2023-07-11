package deamon

import (
	"context"
	"log"

	hubrpc "github.com/Mihalic2040/Hub-rpc"
	"github.com/Mihalic2040/Hub-rpc/src/proto/api"

	"github.com/Mihalic2040/Hub-Mobile/protocols"
)

func (s *Service) Peers(ctx context.Context, req *protocols.PeersResponse) (*protocols.PeersResponse, error) {

	// Implement the logic to handle the gRPC request and return a response
	return &protocols.PeersResponse{Peers: "s.App.Host.Peerstore().Peers().String()"}, nil
}

// request
func (s *Service) Request(ctx context.Context, req *protocols.RequestData) (*protocols.ResponseData, error) {

	data := api.Request{
		Payload: req.Payload,
		User:    req.User,
		Handler: req.Handler,
	}

	log.Println(len(s.host.Peerstore().Peers()))

	resp, err := hubrpc.NewRequest(s.contex, req.Peer, &data, s.cfg.ProtocolId, *s.dht, s.host)
	if err != nil {
		return &protocols.ResponseData{
			Payload: err.Error(),
			Status:  0,
		}, nil
	}

	return &protocols.ResponseData{
		Payload: resp.Payload,
		Status:  resp.Status,
	}, nil

	// return &protocols.ResponseData{
	// 	Payload: "resp.Payload",
	// 	Status:  43,
	// }, nil
}
