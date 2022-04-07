package events

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/quankori/go-eth/internal/connect"
	"github.com/quankori/go-eth/internal/solc/token"
)

func SubscribeEvent(contractAddress string, blockNumber chan uint64) {
	client := connect.ConnectWebsocket()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.TxHash.Hex()) // pointer to event log
			blockNumber <- vLog.BlockNumber
		}
	}
}

func LogsEvent(contractAddress string) {
	client := connect.ConnectWebsocket()
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(16737288),
		ToBlock:   big.NewInt(16737289),
		Addresses: []common.Address{
			common.HexToAddress(contractAddress),
		},
	}

	type LogTransfer struct {
		From   common.Address
		To     common.Address
		Tokens *big.Int
	}

	type LogApproval struct {
		TokenOwner common.Address
		Spender    common.Address
		Tokens     *big.Int
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSig := []byte("Approval(address,address,uint256)")
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			result, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", result[0])
		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")
			var approvalEvent LogApproval
			result, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", result[0])
		}
	}
}
