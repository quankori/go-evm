package main

import (
	"fmt"

	"github.com/quankori/go-eth/internal/file"
	"github.com/quankori/go-eth/pkg/block"
)

func main() {
	data := file.ReadFile()
	blockNumber := make(chan uint64)
	go block.SubscribeNewBlock(blockNumber)
	blockResult := <-blockNumber
	fmt.Println(blockResult)
	blockResultData := block.GetBlockHash(data)
	for _, tx := range blockResultData.Transactions() {
		fmt.Println(tx.Hash().Hex())
	}
}
