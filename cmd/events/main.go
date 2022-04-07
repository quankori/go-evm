package main

import (
	"fmt"

	"github.com/quankori/go-eth/pkg/events"
)

func main() {
	blockNumber := make(chan uint64)
	go events.SubscribeEvent("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", blockNumber)
	blockResult := <-blockNumber
	fmt.Println(blockResult)
	events.LogsEvent("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
}
