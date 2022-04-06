package eth

import (
	"context"
	"log"

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
