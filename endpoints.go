package deamon

import (
	"context"

	"github.com/Mihalic2040/Hub-Mobile/protocols"
	"github.com/libp2p/go-libp2p-core/peer"
)

func (s *Service) Peers(ctx context.Context, req *protocols.PeersResponse) (*protocols.PeersResponse, error) {

	// Implement the logic to handle the gRPC request and return a response
	return &protocols.PeersResponse{Peers: s.App.Host.Peerstore().Peers().String()}, nil
}

// request
func (s *Service) Request(ctx context.Context, req *protocols.RequestData) (*protocols.ResponseData, error) {

	targetPeerID, err := peer.Decode("12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4")
	if err != nil {
		return &protocols.ResponseData{
			Payload: err.Error(),
			Status:  0,
		}, nil
	}

	peer, err := s.App.Dht.FindPeer(context.Background(), targetPeerID)
	if err != nil {
		return &protocols.ResponseData{
			Payload: err.Error(),
			Status:  1,
		}, nil
	}

	err = s.App.Host.Connect(context.Background(), peer)
	if err != nil {
		return &protocols.ResponseData{
			Payload: err.Error(),
			Status:  2,
		}, nil

	}

	return &protocols.ResponseData{
		Payload: "resp.Payload",
		Status:  3,
	}, nil

	// return &protocols.ResponseData{
	// 	Payload: "resp.Payload",
	// 	Status:  43,
	// }, nil
}
