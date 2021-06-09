package testutil

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"killswitch/bridge/abi"
)

func DeployTokenWith(ctx Context, wallet *Wallet, name, symbol string, decimal uint8) (*abi.WrappedToken, common.Address) {
	addr, _, token, err := abi.DeployWrappedToken(wallet.TxOpts, ctx.Backend, name, symbol, decimal)
	if err != nil {
		log.Panicf("can not deploy token; %v", err)
	}
	ctx.Backend.Commit()

	return token, addr
}

func DeployToken(ctx Context, wallet *Wallet) (*abi.WrappedToken, common.Address) {
	return DeployTokenWith(ctx, wallet, "kTest Token", "kTest", 18)
}

func DeployBridgeLocker(ctx Context, wallet *Wallet, token common.Address, name string, fee *big.Int) (*abi.BridgeLocker, common.Address) {
	_, feeAddr := DeployFeeFixed(ctx, wallet, fee)

	addr, _, locker, err := abi.DeployBridgeLocker(wallet.TxOpts, ctx.Backend, token, name, feeAddr, common.BigToAddress(big.NewInt(0)))
	if err != nil {
		log.Panicf("can not deploy bridge locker; %v", err)
	}
	ctx.Backend.Commit()

	return locker, addr
}

func DeployBridgeBurner(ctx Context, wallet *Wallet, token common.Address, name string, fee *big.Int) (*abi.BridgeBurner, common.Address) {
	_, feeAddr := DeployFeeFixed(ctx, wallet, fee)

	addr, _, burner, err := abi.DeployBridgeBurner(wallet.TxOpts, ctx.Backend, token, name, feeAddr, common.BigToAddress(big.NewInt(0)))
	if err != nil {
		log.Panicf("can not deploy burner locker; %v", err)
	}
	ctx.Backend.Commit()

	return burner, addr
}

func DeployBridgeEther(ctx Context, wallet *Wallet, name string, fee *big.Int) (*abi.BridgeEther, common.Address) {
	_, feeAddr := DeployFeeFixed(ctx, wallet, fee)

	addr, _, ether, err := abi.DeployBridgeEther(wallet.TxOpts, ctx.Backend, name, feeAddr, common.BigToAddress(big.NewInt(0)))
	if err != nil {
		log.Panicf("can not deploy bridge locker; %v", err)
	}
	ctx.Backend.Commit()

	return ether, addr
}

func DeployFeeFixed(ctx Context, wallet *Wallet, fee *big.Int) (*abi.FeeFixed, common.Address) {
	feeAddr, _, feeFixed, err := abi.DeployFeeFixed(wallet.TxOpts, ctx.Backend, fee)
	if err != nil {
		log.Panicf("can not deploy fee fixed; %v", err)
	}
	ctx.Backend.Commit()

	return feeFixed, feeAddr
}
