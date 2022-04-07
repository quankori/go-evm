package connect

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quankori/go-eth/config"
)

func ConnectClient() *ethclient.Client {
	config, _ := config.LoadConfig()
	client, err := ethclient.DialContext(context.Background(), config.NetworkURI)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()
	return client
}

func ConnectWebsocket() *ethclient.Client {
	config, _ := config.LoadConfig()
	client, err := ethclient.Dial(config.WebsocketkURI)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
