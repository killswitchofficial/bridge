package abi_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestBridgeEther(t *testing.T) {
	ctx := testutil.Setup(t)

	ether, etherAddr := testutil.DeployBridgeEther(ctx, ctx.Wallets[0], "Test Ether", decimal.EtherToWei("0.1"))

	ownerInitialBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	{
		txOpts := ctx.Wallets[1].TxOpts
		txOpts.Value = decimal.EtherToWei("1.1")

		_, err := ether.Lock(txOpts, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()
	}

	{
		etherBalance := testutil.BalanceETH(ctx, etherAddr)
		require.Equal(t, decimal.EtherToWei("1").String(), etherBalance.String())
	}

	require.Equal(t, new(big.Int).Add(decimal.EtherToWei("0.1"), ownerInitialBalance).String(), testutil.BalanceETH(ctx, ctx.Wallets[0].Address).String())

	{
		initialBalance := testutil.BalanceETH(ctx, ctx.Wallets[2].Address)

		_, err := ether.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[2].Address, decimal.EtherToWei("0.5"), common.Hash{})
		require.NoError(t, err)
		ctx.Backend.Commit()

		etherBalance := testutil.BalanceETH(ctx, etherAddr)
		require.Equal(t, decimal.EtherToWei("0.5").String(), etherBalance.String())

		balance := testutil.BalanceETH(ctx, ctx.Wallets[2].Address)
		require.Equal(t, initialBalance.Add(initialBalance, decimal.EtherToWei("0.5")).String(), balance.String())
	}
}

func TestBridgeEther_RenounceOwnership(t *testing.T) {
	ctx := testutil.Setup(t)

	ether, etherAddr := testutil.DeployBridgeEther(ctx, ctx.Wallets[0], "Test Ether", decimal.EtherToWei("0"))

	{
		txOpts := ctx.Wallets[1].TxOpts
		txOpts.Value = decimal.EtherToWei("1")

		_, err := ether.Lock(txOpts, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()
	}

	{
		balance := testutil.BalanceETH(ctx, etherAddr)
		require.Equal(t, decimal.EtherToWei("1").String(), balance.String())
	}

	initialBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	_, err := ether.RenounceOwnership(ctx.Wallets[0].TxOpts)
	require.NoError(t, err)
	ctx.Backend.Commit()

	{
		balance := testutil.BalanceETH(ctx, etherAddr)
		require.Equal(t, decimal.EtherToWei("0").String(), balance.String())
	}

	paused, err := ether.Paused(nil)
	require.NoError(t, err)
	require.True(t, paused)

	owner, err := ether.Owner(nil)
	require.NoError(t, err)
	require.EqualValues(t, common.Address{}, owner)

	endBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)
	require.NoError(t, err)

	require.Equal(t, 1, endBalance.Cmp(initialBalance))
}

func TestBridgeEther_Receive(t *testing.T) {
	ctx := testutil.Setup(t)

	locker, _ := testutil.DeployBridgeEther(ctx, ctx.Wallets[0], "Test Ether", decimal.EtherToWei("0"))

	_, err := locker.Receive(ctx.Wallets[0].TxOpts)
	require.Error(t, err)

	initialBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	txOpts := ctx.Wallets[0].TxOpts
	txOpts.Value = decimal.EtherToWei("1")
	_, err = locker.Receive(txOpts)
	require.Error(t, err)

	endBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	require.Equal(t, -1, initialBalance.Sub(initialBalance, endBalance).Cmp(decimal.EtherToWei("1")))
}
