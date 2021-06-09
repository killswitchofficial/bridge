package testutil

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"killswitch/bridge/decimal"
)

var chainID = big.NewInt(1337)

// Context is the test context
// provides isolate backend, so each test can run parallel
// without interfere with other
type Context struct {
	context.Context
	Backend *backends.SimulatedBackend
	Wallets []*Wallet
}

// Wallet contains all useful struct for testing
type Wallet struct {
	Key     *ecdsa.PrivateKey
	Address common.Address
	TxOpts  *bind.TransactOpts
}

func Setup(t *testing.T) Context {
	t.Helper()

	var ctx Context
	ctx.Context = context.Background()

	for i := 0; i < 20; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)

		txOpts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
		require.NoError(t, err)

		ctx.Wallets = append(ctx.Wallets, &Wallet{
			Key:     key,
			Address: crypto.PubkeyToAddress(key.PublicKey),
			TxOpts:  txOpts,
		})
	}

	alloc := core.GenesisAlloc{}
	for _, w := range ctx.Wallets {
		alloc[w.TxOpts.From] = core.GenesisAccount{
			Balance: decimal.EtherToWei("10000"),
		}
	}
	ctx.Backend = backends.NewSimulatedBackend(alloc, 10000000)

	return ctx
}

func ExtractSender(tx *types.Transaction) common.Address {
	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
	if err != nil {
		log.Panicf("can not extract sender from tx; %v", err)
	}

	return msg.From()
}

func BalanceETH(ctx Context, address common.Address) *big.Int {
	balance, err := ctx.Backend.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Panicf("can not get balance eth; %v", err)
	}

	return balance
}
