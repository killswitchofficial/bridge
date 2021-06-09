package abi_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"killswitch/bridge/abi"
	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestBridgeBurner(t *testing.T) {
	t.Run("Fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
		burner, _ := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0.1"))

		t.Run("Check Fee", func(t *testing.T) {
			fee, err := burner.CalculateFee(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.1").String(), fee.String())
		})

		t.Run("SetFee", func(t *testing.T) {
			newFeeFixed, newFeeAddr := testutil.DeployFeeFixed(ctx, ctx.Wallets[0], decimal.EtherToWei("0.2"))
			ctx.Backend.Commit()

			_, err := burner.SetFee(ctx.Wallets[0].TxOpts, newFeeAddr)
			require.NoError(t, err)
			ctx.Backend.Commit()

			fee, err := burner.CalculateFee(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.2").String(), fee.String())

			fee, err = newFeeFixed.Calculate(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.2").String(), fee.String())
		})

		t.Run("SetFee not owner", func(t *testing.T) {
			_, newFeeAddr := testutil.DeployFeeFixed(ctx, ctx.Wallets[1], decimal.EtherToWei("0.5"))

			_, err := burner.SetFee(ctx.Wallets[1].TxOpts, newFeeAddr)
			require.Error(t, err)
			ctx.Backend.Commit()
		})
	})

	t.Run("Ownable", func(t *testing.T) {
		ctx := testutil.Setup(t)

		_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
		burner, _ := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0.1"))

		t.Run("Check Owner", func(t *testing.T) {
			owner, err := burner.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[0].Address.String(), owner.String())
		})

		t.Run("TransferOwnership", func(t *testing.T) {
			_, err := burner.TransferOwnership(ctx.Wallets[0].TxOpts, ctx.Wallets[2].Address)
			require.NoError(t, err)
			ctx.Backend.Commit()

			owner, err := burner.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[2].Address.String(), owner.String())

			_, err = burner.TransferOwnership(ctx.Wallets[2].TxOpts, ctx.Wallets[0].Address)
			require.NoError(t, err)
			ctx.Backend.Commit()

			owner, err = burner.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[0].Address.String(), owner.String())
		})

		t.Run("TransferOwnership not owner", func(t *testing.T) {
			_, err := burner.TransferOwnership(ctx.Wallets[2].TxOpts, ctx.Wallets[2].Address)
			require.Error(t, err)
		})
	})

	t.Run("Pause unpause", func(t *testing.T) {
		ctx := testutil.Setup(t)

		// addr0 => bridge owner
		// addr10 => token owner
		// addr1 => user
		userTxOpts := ctx.Wallets[1].TxOpts

		token, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[10], "kBNB", "kBNB", 18)
		_, err := token.AddMinter(ctx.Wallets[10].TxOpts, ctx.Wallets[10].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		burner, burnerAddr := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, burnerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve burner allowance
		_, err = token.Approve(userTxOpts, burnerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// not owner can not pause
		_, err = burner.Pause(ctx.Wallets[10].TxOpts)
		require.Error(t, err)

		// pause
		_, err = burner.Pause(ctx.Wallets[0].TxOpts)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user call to lock wBNB when paused
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.Error(t, err)

		// not owner can not unpause
		_, err = burner.Unpause(ctx.Wallets[10].TxOpts)
		require.Error(t, err)

		// unpause
		_, err = burner.Unpause(ctx.Wallets[0].TxOpts)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user call to lock wBNB when not paused
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
	})

	t.Run("Lock without fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		// addr0 => bridge owner
		// addr10 => token owner
		// addr1 => user
		userTxOpts := ctx.Wallets[1].TxOpts

		token, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[10], "wBNB", "wBNB", 18)
		_, err := token.AddMinter(ctx.Wallets[10].TxOpts, ctx.Wallets[10].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		burner, burnerAddr := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, burnerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve burner allowance
		_, err = token.Approve(userTxOpts, burnerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		initialOwnerBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

		// user call to lock wBNB
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// get locked event
		lockedIt, err := burner.FilterLocked(nil, nil)
		require.NoError(t, err)

		lockedIt.Next()
		lockedIt.Close()

		require.NotNil(t, lockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), lockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.2").String(), lockedIt.Event.Amount.String())

		burnerMustHaveZeroTokenBalance(t, burnerAddr, token)

		// check contract eth balance
		require.Equal(t, decimal.EtherToWei("0").String(), testutil.BalanceETH(ctx, burnerAddr).String())

		// check owner wBNB Balance
		ownerTokenBalance, err := token.BalanceOf(nil, ctx.Wallets[0].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), ownerTokenBalance.String())

		// check owner eth balance
		require.Equal(t, initialOwnerBalance.String(), testutil.BalanceETH(ctx, ctx.Wallets[0].Address).String())
	})

	t.Run("Lock with fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		// addr0 => bridge owner
		// addr10 => token owner
		// addr1 => user
		userTxOpts := ctx.Wallets[1].TxOpts

		token, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[10], "wBNB", "wBNB", 18)
		_, err := token.AddMinter(ctx.Wallets[10].TxOpts, ctx.Wallets[10].Address)
		require.NoError(t, err)

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		burner, burnerAddr := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("1"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, burnerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve burner allowance
		_, err = token.Approve(userTxOpts, burnerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		initialOwnerBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

		// user call to lock wBNB

		// not enough fee
		userTxOpts.Value = decimal.EtherToWei("0.9")
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.Error(t, err)

		// enough fee
		userTxOpts.Value = decimal.EtherToWei("1")
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// reset fund
		userTxOpts.Value = nil

		// get locked event
		lockedIt, err := burner.FilterLocked(nil, nil)
		require.NoError(t, err)

		lockedIt.Next()
		lockedIt.Close()

		require.NotNil(t, lockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), lockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.2").String(), lockedIt.Event.Amount.String())

		burnerMustHaveZeroTokenBalance(t, burnerAddr, token)

		// check contract eth balance
		require.Equal(t, decimal.EtherToWei("0").String(), testutil.BalanceETH(ctx, burnerAddr).String())

		// check owner wBNB Balance
		ownerTokenBalance, err := token.BalanceOf(nil, ctx.Wallets[0].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), ownerTokenBalance.String())

		// check owner eth balance
		finalOwnerBalance := new(big.Int).Add(initialOwnerBalance, decimal.EtherToWei("1"))
		require.Equal(t, finalOwnerBalance.String(), testutil.BalanceETH(ctx, ctx.Wallets[0].Address).String())
	})

	t.Run("Unlock", func(t *testing.T) {
		ctx := testutil.Setup(t)

		// addr0 => bridge owner
		// addr10 => token owner
		// addr1 => user
		userTxOpts := ctx.Wallets[1].TxOpts

		token, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[10], "wBNB", "wBNB", 18)
		_, err := token.AddMinter(ctx.Wallets[10].TxOpts, ctx.Wallets[10].Address)
		require.NoError(t, err)
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		burner, burnerAddr := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, burnerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve burner allowance
		_, err = token.Approve(userTxOpts, burnerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user call to lock wBNB
		_, err = burner.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		burnerMustHaveZeroTokenBalance(t, burnerAddr, token)

		// check user wBNB balance
		userTokenBalance, err := token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.8").String(), userTokenBalance.String())

		// user can not unlock
		_, err = burner.Unlock(userTxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{})
		require.Error(t, err)

		// token owner can not unlock
		_, err = burner.Unlock(ctx.Wallets[1].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{1})
		require.Error(t, err)

		// bridge owner can unlock
		_, err = burner.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{2})
		require.NoError(t, err)
		ctx.Backend.Commit()

		// get unlocked event
		unlockedIt, err := burner.FilterUnlocked(nil, nil)
		require.NoError(t, err)

		unlockedIt.Next()
		unlockedIt.Close()

		require.NotNil(t, unlockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), unlockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.1").String(), unlockedIt.Event.Amount.String())

		burnerMustHaveZeroTokenBalance(t, burnerAddr, token)

		// check user wBNB balance
		userTokenBalance, err = token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.9").String(), userTokenBalance.String())

		// unwrap all
		_, err = burner.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{3})
		require.NoError(t, err)
		ctx.Backend.Commit()

		burnerMustHaveZeroTokenBalance(t, burnerAddr, token)

		// check user wBNB balance
		userTokenBalance, err = token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("1").String(), userTokenBalance.String())
	})
}

func TestBridgeBurner_RenounceOwnership(t *testing.T) {
	ctx := testutil.Setup(t)

	_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
	burner, _ := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0"))

	_, err := burner.RenounceOwnership(ctx.Wallets[0].TxOpts)
	require.NoError(t, err)
	ctx.Backend.Commit()

	paused, err := burner.Paused(nil)
	require.NoError(t, err)
	require.True(t, paused)

	owner, err := burner.Owner(nil)
	require.NoError(t, err)
	require.EqualValues(t, common.Address{}, owner)
}

func TestBridgeBurner_Receive(t *testing.T) {
	ctx := testutil.Setup(t)

	_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
	burner, _ := testutil.DeployBridgeBurner(ctx, ctx.Wallets[0], tokenAddr, "Test Burner", decimal.EtherToWei("0"))

	_, err := burner.Receive(ctx.Wallets[0].TxOpts)
	require.Error(t, err)

	initialBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	txOpts := ctx.Wallets[0].TxOpts
	txOpts.Value = decimal.EtherToWei("1")
	_, err = burner.Receive(txOpts)
	require.Error(t, err)

	endBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

	require.Equal(t, -1, initialBalance.Sub(initialBalance, endBalance).Cmp(decimal.EtherToWei("1")))
}

func burnerMustHaveZeroTokenBalance(t *testing.T, burnerAddr common.Address, token *abi.WrappedToken) {
	balance, err := token.BalanceOf(nil, burnerAddr)
	require.NoError(t, err)
	require.Equal(t, decimal.EtherToWei("0").String(), balance.String())
}
