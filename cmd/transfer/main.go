package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/quankori/go-eth/internal/contracts"
	"github.com/quankori/go-eth/pkg/transaction"
	"github.com/quankori/go-eth/pkg/wallet"
	"github.com/quankori/go-eth/utils/convert"
)

func main() {
	token := contracts.TokenContract()
	result, _ := token.Transfer(wallet.GetAuth(), common.HexToAddress("0xAE628ab2Bf86871D53A32C83759c30AbF3b6B478"), convert.EtherToWei("0.01"))
	receipt := transaction.TransactionHash(*result).Receipt()
	fmt.Println(receipt.TxHash.String())
}
