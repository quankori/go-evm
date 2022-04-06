package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/quankori/go-eth/config"
	"github.com/quankori/go-eth/internal/connect"
	"github.com/quankori/go-eth/pkg/contracts"
)

func main() {

	// block := eth.GetCurrentBlock()
	// fmt.Println(block.Number())
	// balance := eth.GetBalance("your_public_key").ParseEther()
	// fmt.Println(balance)
	// info := wallet.GenerateWallet()
	// fmt.Println(info.PrivateKey)
	// transaction.Transfer(common.HexToAddress("your_wallet"), big.NewInt(10000000000000000))
	client := connect.ConnectClient()
	config, _ := config.LoadConfig()
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	token := contracts.TokenContract()
	result, _ := token.Transfer(auth, common.HexToAddress("0xAE628ab2Bf86871D53A32C83759c30AbF3b6B478"), big.NewInt(100000000000000000))
	fmt.Println(result.Hash().String())
	tx, isPending, err := client.TransactionByHash(context.Background(), result.Hash())
	fmt.Println(isPending)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	fmt.Println(receipt.Status)
	fmt.Println(receipt.BlockNumber)
}
