package abi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"killswitch/bridge/abi"
	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestLimiter(t *testing.T) {
	ctx := testutil.Setup(t)

	_, _, limit, err := abi.DeployLimiterDaily(ctx.Wallets[0].TxOpts, ctx.Backend)
	require.NoError(t, err)
	ctx.Backend.Commit()

	t.Run("Empty", func(t *testing.T) {
		addr := ctx.Wallets[10].Address

		usage, err := limit.GetUsage(nil, addr)
		require.NoError(t, err)
		require.Equal(t, "0", usage.String())

		isLimited, err := limit.IsLimited(nil, addr, decimal.EtherToWei("1"))
		require.NoError(t, err)
		require.False(t, isLimited)

		limitValue, err := limit.GetLimit(nil, addr)
		require.NoError(t, err)
		require.Equal(t, "0", limitValue.String())

		_, err = limit.IncreaseUsage(ctx.Wallets[10].TxOpts, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()
	})

	t.Run("setLimiter not owner", func(t *testing.T) {
		_, err := limit.SetLimit(ctx.Wallets[1].TxOpts, ctx.Wallets[10].Address, decimal.EtherToWei("1"))
		require.Error(t, err)
	})

	t.Run("setLimiter", func(t *testing.T) {
		addr1 := ctx.Wallets[11].Address
		addr2 := ctx.Wallets[12].Address

		_, err := limit.SetLimit(ctx.Wallets[0].TxOpts, addr1, decimal.EtherToWei("1"))
		require.NoError(t, err)
		_, err = limit.SetLimit(ctx.Wallets[0].TxOpts, addr2, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		limitValue, err := limit.GetLimit(nil, addr1)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("1").String(), limitValue.String())

		usage, err := limit.GetUsage(nil, addr1)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), usage.String())

		isLimited, err := limit.IsLimited(nil, addr1, decimal.EtherToWei("0.5"))
		require.NoError(t, err)
		require.False(t, isLimited)

		isLimited, err = limit.IsLimited(nil, addr1, decimal.EtherToWei("1"))
		require.NoError(t, err)
		require.False(t, isLimited)

		isLimited, err = limit.IsLimited(nil, addr1, decimal.EtherToWei("1.1"))
		require.NoError(t, err)
		require.True(t, isLimited)

		_, err = limit.IncreaseUsage(ctx.Wallets[11].TxOpts, decimal.EtherToWei("0.5"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		usage, err = limit.GetUsage(nil, addr1)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.5").String(), usage.String())

		usage, err = limit.GetUsage(nil, addr2)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), usage.String())

		_, err = limit.IncreaseUsage(ctx.Wallets[11].TxOpts, decimal.EtherToWei("0.6"))
		require.Error(t, err)

		_, err = limit.IncreaseUsage(ctx.Wallets[11].TxOpts, decimal.EtherToWei("0.5"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		usage, err = limit.GetUsage(nil, addr1)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("1").String(), usage.String())

		usage, err = limit.GetUsage(nil, addr2)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), usage.String())

		isLimited, err = limit.IsLimited(nil, addr1, decimal.EtherToWei("0"))
		require.NoError(t, err)
		require.False(t, isLimited)

		isLimited, err = limit.IsLimited(nil, addr1, decimal.EtherToWei("0.01"))
		require.NoError(t, err)
		require.True(t, isLimited)

		_, err = limit.IncreaseUsage(ctx.Wallets[11].TxOpts, decimal.EtherToWei("0"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = limit.IncreaseUsage(ctx.Wallets[11].TxOpts, decimal.EtherToWei("0.01"))
		require.Error(t, err)
	})
}
