package eth

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/quankori/go-eth/internal/connect"
	"github.com/quankori/go-eth/utils/convert"
)

// Get balance user
func GetBalance(walletAddress string) *convert.BigInteger {
	client := connect.ConnectClient()
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(walletAddress), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	return convert.NewBigInteger(balance)
}
