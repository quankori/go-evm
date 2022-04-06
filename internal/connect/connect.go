package connect

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var initURL = "https://data-seed-prebsc-1-s1.binance.org:8545"

func ConnectClient() *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), initURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()
	return client
}
