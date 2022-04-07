package wallet

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/quankori/go-eth/config"
	"github.com/quankori/go-eth/internal/connect"
)

// Get balance user
func GetBalance(walletAddress string) *big.Int {
	client := connect.ConnectClient()
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(walletAddress), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	return balance
}

func GetAuth() *bind.TransactOpts {
	client := connect.ConnectClient()
	config, _ := config.LoadConfig()
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	return auth
}
