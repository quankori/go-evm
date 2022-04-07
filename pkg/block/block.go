package block

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/quankori/go-eth/internal/connect"
)

func GetCurrentBlock() *types.Block {
	client := connect.ConnectClient()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	return block
}

func SubscribeNewBlock(blockNumber chan uint64) uint64 {
	client := connect.ConnectWebsocket()
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			blockNumber <- block.Number().Uint64()
		}
	}
}

func GetBlockHash(blockNumber int64) *types.Block {
	client := connect.ConnectWebsocket()
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}
	return block
}
