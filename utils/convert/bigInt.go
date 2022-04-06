package convert

import (
	"math"
	"math/big"
)

type BigInteger struct {
	i *big.Int
}

func NewBigInteger(c *big.Int) *BigInteger {
	return &BigInteger{c}
}

// Parse wei to ether
func (balance BigInteger) ParseEther() *big.Float {
	fBalance := new(big.Float)
	fBalance.SetString(balance.i.String())
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return balanceEther
}
