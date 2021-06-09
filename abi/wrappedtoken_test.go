package abi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestWrappedToken(t *testing.T) {
	ctx := testutil.Setup(t)

	token, _ := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "Test Token", "TEST", 18)

	// metadata
	{
		name, err := token.Name(nil)
		require.NoError(t, err)
		require.Equal(t, "Test Token", name)

		symbol, err := token.Symbol(nil)
		require.NoError(t, err)
		require.Equal(t, "TEST", symbol)

		decimals, err := token.Decimals(nil)
		require.NoError(t, err)
		require.EqualValues(t, 18, decimals)
	}

	// owner can not mint
	{
		_, err := token.Mint(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1.5"))
		require.Error(t, err)
	}

	// mint
	{
		// add owner to be minter
		_, err := token.AddMinter(ctx.Wallets[0].TxOpts, ctx.Wallets[0].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1.5"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		balance, err := token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1.5"), balance)

		totalSupply, err := token.TotalSupply(nil)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1.5"), totalSupply)
	}

	// not minter can not mint
	{
		_, err := token.Mint(ctx.Wallets[1].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1.5"))
		require.Error(t, err)
	}

	// remove owner from minter then should not able to mint
	{
		// remove owner to be minter
		_, err := token.RemoveMinter(ctx.Wallets[0].TxOpts, ctx.Wallets[0].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1.5"))
		require.Error(t, err)
	}

	// transfer
	{
		_, err := token.Transfer(ctx.Wallets[1].TxOpts, ctx.Wallets[2].Address, decimal.EtherToWei("0.5"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		balance, err := token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1"), balance)

		balance, err = token.BalanceOf(nil, ctx.Wallets[2].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("0.5"), balance)

		totalSupply, err := token.TotalSupply(nil)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1.5"), totalSupply)
	}

	// not minter can not burn
	{
		_, err := token.Burn(ctx.Wallets[2].TxOpts, decimal.EtherToWei("0.1"))
		require.Error(t, err)
	}

	// burn
	{
		// not minter can not burn
		_, err := token.Burn(ctx.Wallets[2].TxOpts, decimal.EtherToWei("0.1"))
		require.Error(t, err)

		// add owner to minter
		_, err = token.AddMinter(ctx.Wallets[0].TxOpts, ctx.Wallets[0].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[0].TxOpts, ctx.Wallets[0].Address, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Burn(ctx.Wallets[0].TxOpts, decimal.EtherToWei("0.1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		balance, err := token.BalanceOf(nil, ctx.Wallets[0].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("0.1"), balance)

		balance, err = token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1"), balance)

		balance, err = token.BalanceOf(nil, ctx.Wallets[2].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("0.5"), balance)

		totalSupply, err := token.TotalSupply(nil)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("1.6"), totalSupply)
	}

	// burnFrom
	{
		_, err := token.Mint(ctx.Wallets[0].TxOpts, ctx.Wallets[15].Address, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// not minter can not burn
		_, err = token.BurnFrom(ctx.Wallets[2].TxOpts, ctx.Wallets[15].Address, decimal.EtherToWei("0.1"))
		require.Error(t, err)
		_, err = token.BurnFrom(ctx.Wallets[15].TxOpts, ctx.Wallets[15].Address, decimal.EtherToWei("0.1"))
		require.Error(t, err)

		// without approve
		_, err = token.BurnFrom(ctx.Wallets[0].TxOpts, ctx.Wallets[15].Address, decimal.EtherToWei("0.1"))
		require.Error(t, err)

		// approve
		_, err = token.Approve(ctx.Wallets[15].TxOpts, ctx.Wallets[0].Address, decimal.EtherToWei("0.1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.BurnFrom(ctx.Wallets[0].TxOpts, ctx.Wallets[15].Address, decimal.EtherToWei("0.1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		balance, err := token.BalanceOf(nil, ctx.Wallets[15].Address)
		require.NoError(t, err)
		require.EqualValues(t, decimal.EtherToWei("0.1"), balance)
	}

	// // burn log
	// {
	// 	it, err := token.FilterBurn(nil, nil)
	// 	require.NoError(t, err)
	//
	// 	it.Next()
	// 	it.Close()
	//
	// 	require.NotEmpty(t, it.Event)
	// 	require.EqualValues(t, ctx.Wallets[0].Address, it.Event.Sender)
	// }

	// // burnTo
	// {
	// 	_, err := token.BurnTo(ctx.Wallets[2].TxOpts, testutil.ETHToWei(0.2), ctx.Wallets[3].Address)
	// 	assert.NoError(t, err)
	//
	// 	balance, err := token.BalanceOf(nil, ctx.Wallets[2].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.2), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[3].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0).String(), balance.String())
	//
	// 	it, err := token.FilterBurn(nil, nil, nil)
	// 	assert.NoError(t, err)
	//
	// 	it.Next()
	// 	it.Next()
	// 	it.Close()
	//
	// 	assert.NotEmpty(t, it.Event)
	// 	assert.EqualValues(t, ctx.Wallets[3].Address, it.Event.To)
	// }

	// // mintBatch
	// {
	// 	_, err := token.MintBatch(ctx.Wallets[0].TxOpts,
	// 		[]common.Address{
	// 			ctx.Wallets[4].Address,
	// 			ctx.Wallets[5].Address,
	// 			ctx.Wallets[6].Address,
	// 			ctx.Wallets[7].Address,
	// 		},
	// 		[]*big.Int{
	// 			testutil.ETHToWei(0.1),
	// 			testutil.ETHToWei(0.2),
	// 			testutil.ETHToWei(0.3),
	// 			testutil.ETHToWei(0.4),
	// 		},
	// 	)
	// 	assert.NoError(t, err)
	//
	// 	balance, err := token.BalanceOf(nil, ctx.Wallets[4].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.1), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[5].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.2), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[6].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.3), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[7].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.4), balance)
	// }
	//
	// // mintBatch empty accounts
	// {
	// 	_, err := token.MintBatch(ctx.Wallets[0].TxOpts, []common.Address{}, []*big.Int{})
	// 	assert.Error(t, err)
	// }
	//
	// // mintBatch not equal accounts and amounts
	// {
	// 	_, err := token.MintBatch(ctx.Wallets[0].TxOpts,
	// 		[]common.Address{
	// 			ctx.Wallets[4].Address,
	// 			ctx.Wallets[5].Address,
	// 			ctx.Wallets[6].Address,
	// 			ctx.Wallets[7].Address,
	// 		},
	// 		[]*big.Int{
	// 			testutil.ETHToWei(5),
	// 			testutil.ETHToWei(5),
	// 			testutil.ETHToWei(5),
	// 		},
	// 	)
	// 	assert.Error(t, err)
	//
	// 	balance, err := token.BalanceOf(nil, ctx.Wallets[4].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.1), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[5].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.2), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[6].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.3), balance)
	//
	// 	balance, err = token.BalanceOf(nil, ctx.Wallets[7].Address)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, testutil.ETHToWei(0.4), balance)
	// }
	//
	// // not owner can not mintBatch
	// {
	// 	_, err := token.MintBatch(ctx.Wallets[1].TxOpts, []common.Address{ctx.Wallets[1].Address}, []*big.Int{testutil.ETHToWei(1.5)})
	// 	assert.Error(t, err)
	// }
}
