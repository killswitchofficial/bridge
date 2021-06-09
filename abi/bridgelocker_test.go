package abi_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"killswitch/bridge/decimal"
	"killswitch/bridge/testutil"
)

func TestBridgeLocker(t *testing.T) {
	t.Run("Fee", func(t *testing.T) {
		ctx := testutil.Setup(t)

		_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
		locker, _ := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0.1"))

		t.Run("Check Fee", func(t *testing.T) {
			fee, err := locker.CalculateFee(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.1").String(), fee.String())
		})

		t.Run("SetFee", func(t *testing.T) {
			newFeeFixed, newFeeAddr := testutil.DeployFeeFixed(ctx, ctx.Wallets[0], decimal.EtherToWei("0.2"))
			ctx.Backend.Commit()

			_, err := locker.SetFee(ctx.Wallets[0].TxOpts, newFeeAddr)
			require.NoError(t, err)
			ctx.Backend.Commit()

			fee, err := locker.CalculateFee(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.2").String(), fee.String())

			fee, err = newFeeFixed.Calculate(nil, decimal.EtherToWei("0"))
			require.NoError(t, err)
			require.Equal(t, decimal.EtherToWei("0.2").String(), fee.String())
		})

		t.Run("SetFee not owner", func(t *testing.T) {
			_, newFeeAddr := testutil.DeployFeeFixed(ctx, ctx.Wallets[1], decimal.EtherToWei("0.5"))

			_, err := locker.SetFee(ctx.Wallets[1].TxOpts, newFeeAddr)
			require.Error(t, err)
		})
	})

	t.Run("Ownable", func(t *testing.T) {
		ctx := testutil.Setup(t)

		_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
		locker, _ := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0.1"))

		t.Run("Check Owner", func(t *testing.T) {
			owner, err := locker.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[0].Address.String(), owner.String())
		})

		t.Run("TransferOwnership", func(t *testing.T) {
			_, err := locker.TransferOwnership(ctx.Wallets[0].TxOpts, ctx.Wallets[2].Address)
			require.NoError(t, err)
			ctx.Backend.Commit()

			owner, err := locker.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[2].Address.String(), owner.String())

			_, err = locker.TransferOwnership(ctx.Wallets[2].TxOpts, ctx.Wallets[0].Address)
			require.NoError(t, err)
			ctx.Backend.Commit()

			owner, err = locker.Owner(nil)
			require.NoError(t, err)
			require.Equal(t, ctx.Wallets[0].Address.String(), owner.String())
		})

		t.Run("TransferOwnership not owner", func(t *testing.T) {
			_, err := locker.TransferOwnership(ctx.Wallets[2].TxOpts, ctx.Wallets[2].Address)
			require.Error(t, err)
		})
	})

	t.Run("Pause unpause", func(t *testing.T) {
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

		locker, lockerAddr := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, lockerAddr)
		require.NoError(t, err)

		// user approve locker allowance
		_, err = token.Approve(userTxOpts, lockerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)

		// not owner can not pause
		_, err = locker.Pause(ctx.Wallets[10].TxOpts)
		require.Error(t, err)

		// pause
		_, err = locker.Pause(ctx.Wallets[0].TxOpts)
		require.NoError(t, err)

		// user call to lock wBNB when paused
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.Error(t, err)

		// not owner can not unpause
		_, err = locker.Unpause(ctx.Wallets[10].TxOpts)
		require.Error(t, err)

		// unpause
		_, err = locker.Unpause(ctx.Wallets[0].TxOpts)
		require.NoError(t, err)

		// user call to lock wBNB when not paused
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
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

		locker, lockerAddr := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, lockerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve locker allowance
		_, err = token.Approve(userTxOpts, lockerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		initialOwnerBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

		// user call to lock wBNB
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// get locked event
		lockedIt, err := locker.FilterLocked(nil, nil)
		require.NoError(t, err)

		lockedIt.Next()
		lockedIt.Close()

		require.NotNil(t, lockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), lockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.2").String(), lockedIt.Event.Amount.String())

		// check contract wBNB balance
		contractTokenBalance, err := token.BalanceOf(nil, lockerAddr)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.2").String(), contractTokenBalance.String())

		// check contract eth balance
		require.Equal(t, decimal.EtherToWei("0").String(), testutil.BalanceETH(ctx, lockerAddr).String())

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
		ctx.Backend.Commit()

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		locker, lockerAddr := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("1"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, lockerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve locker allowance
		_, err = token.Approve(userTxOpts, lockerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		initialOwnerBalance := testutil.BalanceETH(ctx, ctx.Wallets[0].Address)

		// user call to lock wBNB

		// not enough fee
		userTxOpts.Value = decimal.EtherToWei("0.9")
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.Error(t, err)

		// enough fee
		userTxOpts.Value = decimal.EtherToWei("1")
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// reset fund
		userTxOpts.Value = nil

		// get locked event
		lockedIt, err := locker.FilterLocked(nil, nil)
		require.NoError(t, err)

		lockedIt.Next()
		lockedIt.Close()

		require.NotNil(t, lockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), lockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.2").String(), lockedIt.Event.Amount.String())

		// check contract wBNB balance
		contractTokenBalance, err := token.BalanceOf(nil, lockerAddr)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.2").String(), contractTokenBalance.String())

		// check contract eth balance
		require.Equal(t, decimal.EtherToWei("0").String(), testutil.BalanceETH(ctx, lockerAddr).String())

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

		_, err = token.Mint(ctx.Wallets[10].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("1"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		locker, lockerAddr := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0"))

		_, err = token.AddMinter(ctx.Wallets[10].TxOpts, lockerAddr)
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user approve locker allowance
		_, err = token.Approve(userTxOpts, lockerAddr, decimal.EtherToWei("100"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// user call to lock wBNB
		_, err = locker.Lock(userTxOpts, decimal.EtherToWei("0.2"))
		require.NoError(t, err)
		ctx.Backend.Commit()

		// check contract wBNB balance
		contractTokenBalance, err := token.BalanceOf(nil, lockerAddr)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.2").String(), contractTokenBalance.String())

		// check user wBNB balance
		userTokenBalance, err := token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.8").String(), userTokenBalance.String())

		// user can not unlock
		_, err = locker.Unlock(userTxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{})
		require.Error(t, err)

		// token owner can not unlock
		_, err = locker.Unlock(ctx.Wallets[1].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{1})
		require.Error(t, err)

		// bridge owner can unlock
		_, err = locker.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{2})
		require.NoError(t, err)
		ctx.Backend.Commit()

		// get unlocked event
		unlockedIt, err := locker.FilterUnlocked(nil, nil)
		require.NoError(t, err)

		unlockedIt.Next()
		unlockedIt.Close()

		require.NotNil(t, unlockedIt.Event)
		require.Equal(t, ctx.Wallets[1].Address.String(), unlockedIt.Event.Sender.String())
		require.Equal(t, decimal.EtherToWei("0.1").String(), unlockedIt.Event.Amount.String())

		// check contract wBNB balance
		contractTokenBalance, err = token.BalanceOf(nil, lockerAddr)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.1").String(), contractTokenBalance.String())

		// check user wBNB balance
		userTokenBalance, err = token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0.9").String(), userTokenBalance.String())

		// unwrap all
		_, err = locker.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{3})
		require.NoError(t, err)
		ctx.Backend.Commit()

		// check contract wBNB balance
		contractTokenBalance, err = token.BalanceOf(nil, lockerAddr)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("0").String(), contractTokenBalance.String())

		// check user wBNB balance
		userTokenBalance, err = token.BalanceOf(nil, ctx.Wallets[1].Address)
		require.NoError(t, err)
		require.Equal(t, decimal.EtherToWei("1").String(), userTokenBalance.String())

		// unwrap empty contract
		_, err = locker.Unlock(ctx.Wallets[0].TxOpts, ctx.Wallets[1].Address, decimal.EtherToWei("0.1"), [32]byte{4})
		require.Error(t, err)
	})
}

func TestBridgeLocker_RenounceOwnership(t *testing.T) {
	ctx := testutil.Setup(t)

	token, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
	locker, lockerAddr := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0"))

	_, err := token.AddMinter(ctx.Wallets[0].TxOpts, ctx.Wallets[0].Address)
	require.NoError(t, err)
	ctx.Backend.Commit()

	_, err = token.Mint(ctx.Wallets[0].TxOpts, lockerAddr, decimal.EtherToWei("1"))
	require.NoError(t, err)
	ctx.Backend.Commit()

	balance, err := token.BalanceOf(nil, ctx.Wallets[0].Address)
	require.NoError(t, err)
	require.Equal(t, decimal.EtherToWei("0").String(), balance.String())

	_, err = locker.RenounceOwnership(ctx.Wallets[0].TxOpts)
	require.NoError(t, err)
	ctx.Backend.Commit()

	paused, err := locker.Paused(nil)
	require.NoError(t, err)
	require.True(t, paused)

	owner, err := locker.Owner(nil)
	require.NoError(t, err)
	require.EqualValues(t, common.Address{}, owner)

	balance, err = token.BalanceOf(nil, ctx.Wallets[0].Address)
	require.NoError(t, err)

	require.Equal(t, decimal.EtherToWei("1").String(), balance.String())
}

func TestBridgeLocker_Receive(t *testing.T) {
	ctx := testutil.Setup(t)

	_, tokenAddr := testutil.DeployTokenWith(ctx, ctx.Wallets[0], "wBNB", "wBNB", 18)
	locker, _ := testutil.DeployBridgeLocker(ctx, ctx.Wallets[0], tokenAddr, "Test Locker", decimal.EtherToWei("0"))

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
