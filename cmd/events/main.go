package main

import (
	"fmt"

	"github.com/quankori/go-eth/pkg/eth"
	"github.com/quankori/go-eth/pkg/wallet"
)

func main() {
	block := eth.GetCurrentBlock()
	fmt.Println(block.Number())
	balance := eth.GetBalance("0x042F9D3b08D38D388F01192383e962d58A5F99e6").ParseEther()
	fmt.Println(balance)
	info := wallet.GenerateWallet()
	fmt.Println(info.PrivateKey)
}
