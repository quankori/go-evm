package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/quankori/go-eth/pkg/transaction"
)

func main() {

	// block := eth.GetCurrentBlock()
	// fmt.Println(block.Number())
	// balance := eth.GetBalance("your_public_key").ParseEther()
	// fmt.Println(balance)
	// info := wallet.GenerateWallet()
	// fmt.Println(info.PrivateKey)

	transaction.Transfer(common.HexToAddress("your_wallet"), big.NewInt(10000000000000000))
}
