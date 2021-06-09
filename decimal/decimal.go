package decimal

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
)

// EtherToWei converts string ether to wei
func EtherToWei(eth string) *big.Int {
	d, _ := decimal.NewFromString(eth)
	d = d.Mul(decimal.NewFromFloat(params.Ether))
	return d.BigInt()
}

// WeiToEther converts wei to estimate ether (rounding)
func WeiToEther(v *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(v), big.NewFloat(params.Ether))
}
