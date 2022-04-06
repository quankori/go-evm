package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/quankori/go-eth/config"
	"github.com/quankori/go-eth/internal/connect"
)

func Transfer(to common.Address, amount *big.Int) {
	config, _ := config.LoadConfig()
	client := connect.ConnectClient()
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	from := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
