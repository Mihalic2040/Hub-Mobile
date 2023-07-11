package deamon

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	hubrpc "github.com/Mihalic2040/Hub-rpc"
	"github.com/fatih/color"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/mplex"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
	"github.com/multiformats/go-multiaddr"
)

func CreateInstance(ctx context.Context, cfg Config) *host.Host {

	//parese addresses

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/", cfg.Host, cfg.Port))
	sourceMultiAddrQuic, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/%s/quic", cfg.Host, cfg.Port))
	var sourceMultiAddrWs multiaddr.Multiaddr
	if cfg.Port != "0" {
		port, _ := strconv.Atoi(cfg.Port)
		portNumber := strconv.Itoa(port + 1)
		sourceMultiAddrWs, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/ws", cfg.Host, portNumber))
	} else {
		sourceMultiAddrWs, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/ws", cfg.Host, cfg.Port))
	}

	taranspors := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(quic.NewTransport),

		//For js nodes
		libp2p.Transport(websocket.New),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/mplex/", mplex.DefaultTransport),

		//For js nodes
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
	)

	addrs := libp2p.ListenAddrStrings(
		sourceMultiAddr.String(),
		sourceMultiAddrQuic.String(),
		sourceMultiAddrWs.String(),
	)

	host, err := libp2p.New(
		muxers,
		taranspors,
		addrs,
	)
	if err != nil {
		log.Panic("Err creating host", err)
	}

	host.SetStreamHandler(protocol.ID(cfg.ProtocolId), func(s network.Stream) {
		hubrpc.Stream_handler(s, hubrpc.HandlerMap{})
	})

	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("[*] Your ID is: %s", green(host.ID().Pretty()))
	log.Println("[*] Your Multiaddress is:", host.Addrs())

	return &host

}

func CreateDht(ctx context.Context, host host.Host, config Config) *dht.IpfsDHT {
	// init DHT

	if config.DHTServer {
		kademliaDHT, err := dht.New(ctx, host, dht.Mode(dht.ModeAutoServer))
		if err != nil {
			log.Println("[DHT] Fail to init DHT: ", err)
		}
		log.Println("[DHT:Server MODE] Init sucsesfull")
		return kademliaDHT
	} else {
		kademliaDHT, err := dht.New(ctx, host)
		if err != nil {
			log.Println("[DHT] Fail to init DHT: ", err)
		}
		log.Println("[DHT] Init sucsesfull")
		return kademliaDHT
	}

}

func Bootstrap(ctx context.Context, dht *dht.IpfsDHT) {
	//Bootstrap
	log.Println("[DHT] Running bootstrap thread")
	if err := dht.Bootstrap(ctx); err != nil {
		log.Println("[DHT] Fail to bootstrap DHT: ", err)
	}

}

func Boot(ctx context.Context, config Config, host host.Host) {
	if config.Bootstrap == "" {
		return
	}
	log.Println("[DHT:Bootstrap] Running bootstrap from config: ", config.Bootstrap)
	// Parse configuration
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf(config.Bootstrap))
	if err != nil {
		log.Println("[DHT:Bootstrap] Fail to parse multiaddr: ", err)
	}

	// Conver the multiaddr to peerinfo
	var wg sync.WaitGroup
	peerinfo, err := peer.AddrInfoFromP2pAddr(sourceMultiAddr)
	if err != nil {
		log.Println("[DHT:Bootstrap] Fail to parse multiaddr: ", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := host.Connect(ctx, *peerinfo); err != nil {
			log.Println("[DHT:Bootstrap] Fail to connect:", err)
		} else {
			log.Println("[DHT:Bootstrap] Successfully connected to node: ", *peerinfo)
		}
	}()
	wg.Wait()

}

func DiscoverPeers(ctx context.Context, h host.Host) {

	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("Bootstrap warning:", err)
			}
		}()
	}
	wg.Wait()
}
