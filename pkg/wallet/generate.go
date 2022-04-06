package wallet

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey    string
	PublicKey     string
	PublicAddress string
}

func GenerateWallet() *Wallet {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pData := hexutil.Encode(crypto.FromECDSA(pvk))
	puData := hexutil.Encode(crypto.FromECDSAPub(&pvk.PublicKey))
	puAddress := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	return &Wallet{pData, puData, puAddress}
}
