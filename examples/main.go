package main

import (
	"time"

	deamon "github.com/Mihalic2040/Hub-Mobile"
)

func main() {
	service := deamon.Service{}

	configJson := `{
  "Host": "0.0.0.0",
  "Port": "0",
  "Rendezvous": "Hub",
  "DHTServer": true,
  "ProtocolId": "/hub/0.0.1",
  "Bootstrap": "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE"
}`

	service.Config(configJson)
	service.Start()

	time.Sleep(time.Second * 1)

}
