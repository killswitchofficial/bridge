package abi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestFeeFixed(t *testing.T) {
	t.Run("Zero fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		feeFixed, _ := testutil.DeployFeeFixed(ctx, ctx.Wallets[0], decimal.EtherToWei("0"))

		zero := decimal.EtherToWei("0")

		fee, err := feeFixed.Fee(nil)
		require.NoError(t, err)
		require.Equal(t, zero.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("0"))
		require.NoError(t, err)
		require.Equal(t, zero.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("1"))
		require.NoError(t, err)
		require.Equal(t, zero.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("10000"))
		require.NoError(t, err)
		require.Equal(t, zero.String(), fee.String())
	})

	t.Run("Fixed fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		feeFixed, _ := testutil.DeployFeeFixed(ctx, ctx.Wallets[0], decimal.EtherToWei("0.1"))

		value := decimal.EtherToWei("0.1")

		fee, err := feeFixed.Fee(nil)
		require.NoError(t, err)
		require.Equal(t, value.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("0"))
		require.NoError(t, err)
		require.Equal(t, value.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("1"))
		require.NoError(t, err)
		require.Equal(t, value.String(), fee.String())

		fee, err = feeFixed.Calculate(nil, decimal.EtherToWei("10000"))
		require.NoError(t, err)
		require.Equal(t, value.String(), fee.String())
	})

	t.Run("SetFee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		feeFixed, _ := testutil.DeployFeeFixed(ctx, ctx.Wallets[0], decimal.EtherToWei("0.1"))

		// not owner
		_, err := feeFixed.SetFee(ctx.Wallets[1].TxOpts, decimal.EtherToWei("0.2"))
		require.Error(t, err)

		// owner
		_, err = feeFixed.SetFee(ctx.Wallets[0].TxOpts, decimal.EtherToWei("0.3"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		fee, err := feeFixed.Fee(nil)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.3").String(), fee.String())
	})
}
