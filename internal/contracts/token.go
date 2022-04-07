package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/quankori/go-eth/config"
	"github.com/quankori/go-eth/internal/connect"
	"github.com/quankori/go-eth/internal/solc/token"
)

func TokenContract() *token.Token {
	config, _ := config.LoadConfig()
	client := connect.ConnectClient()
	address := common.HexToAddress(config.TokenURI)
	token, err := token.NewToken(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	return token
}
