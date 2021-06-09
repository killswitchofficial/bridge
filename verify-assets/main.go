package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"killswitch/bridge/abi"
)

var bscClient, _ = ethclient.Dial("https://bsc-dataseed.binance.org")
var bkcClient, _ = ethclient.Dial("https://rpc.bitkubchain.io")
var maticClient, _ = ethclient.Dial("https://rpc-mainnet.maticvigil.com")

var client = map[string]*ethclient.Client{
	"bsc":   bscClient,
	"bkc":   bkcClient,
	"matic": maticClient,
}

type ClientWithAddress struct {
	Client  string
	Address string
}

type pairLockBurn struct {
	Name   string
	Locker ClientWithAddress
	Burner ClientWithAddress
}

type pairEtherBurn struct {
	Name   string
	Ether  ClientWithAddress
	Burner ClientWithAddress
}

var verifyPairs = []interface{}{
	// bsc => bkc
	pairEtherBurn{
		"BNB <=> kBNB",
		ClientWithAddress{"bsc", "0xa4e3a7DE03D4138620EEc38766C06d175dF64963"},
		ClientWithAddress{"bkc", "0x87d4E41CA7D2744B95055768F91BdC8B673B7C5E"},
	},
	pairLockBurn{
		"Dolly <=> kDolly",
		ClientWithAddress{"bsc", "0x3bb24415c501Eeaf8b0778C2e306857C89Bb7a23"},
		ClientWithAddress{"bkc", "0xc5C7BbF13Decf2d667bC6287385149E2ba3Eb7D6"},
	},
	pairLockBurn{
		"UST <=> kUST",
		ClientWithAddress{"bsc", "0x8CB22Dd24E930d685e25E5Ec3A4948974e0Cc32c"},
		ClientWithAddress{"bkc", "0x659B98BF5Aa80CBFf74236486915951233169910"},
	},
	pairLockBurn{
		"DAI <=> kDAI",
		ClientWithAddress{"bsc", "0xAA23Db1B0D19f933504c7e2C9279d427834f3692"},
		ClientWithAddress{"bkc", "0xA7E186636Bcb7Da5B6E1aa58aC34DE5D35772d10"},
	},
	pairLockBurn{
		"WMMP <=> kMMP",
		ClientWithAddress{"bsc", "0x3AbE2205740198b651361bAB1E77210D8C247576"},
		ClientWithAddress{"bkc", "0xe79b6ea8C1562e61A184898fB15391a4f538F5D1"},
	},
	pairLockBurn{
		"SZO <=> kSZO",
		ClientWithAddress{"bsc", "0x144F00ef491BB058eA8A56f2B9bFA598a3DfBac6"},
		ClientWithAddress{"bkc", "0x70a0f9Adc1bD39065B48c80BEfd5092814c9bC92"},
	},
	pairLockBurn{
		"CAKE <=> kCAKE",
		ClientWithAddress{"bsc", "0xdB834703FfEA7D0DD173Cf03A7b0a5115dcc03FE"},
		ClientWithAddress{"bkc", "0x7b841f79Adf5d9d475b0501Da9E9092f08eF4cA9"},
	},
	// bkc => bsc
	pairEtherBurn{
		"KUB <=> KUB",
		ClientWithAddress{"bkc", "0x244518458ea1B3f2B0c02C6420Ed160E1ca5c866"},
		ClientWithAddress{"bsc", "0xc0f8Bf1c447c25F52cc7d69f0bBBF8CD5856e66f"},
	},
	pairLockBurn{
		"TUK <=> kTUK",
		ClientWithAddress{"bkc", "0xB70D650d229A4c5Ff67522e69bc38b0E1d9eAAC0"},
		ClientWithAddress{"bsc", "0x6CAa59A946FeEEd92bC923aa15A19539b8988353"},
	},
	// matic => bsc
	pairEtherBurn{
		"MATIC <=> kMATIC",
		ClientWithAddress{"matic", "0x987e283e6B34CCbf069C1d0075f43A12b79142E1"},
		ClientWithAddress{"bsc", "0xED7B8606270295d1b3b60b99c051de4D7D2f7ff2"},
	},
}

func main() {
	ctx := context.Background()

	valid := true

	for _, p := range verifyPairs {
		switch p := p.(type) {
		case pairEtherBurn:
			etherClient := client[p.Ether.Client]

			burnClient := client[p.Burner.Client]
			bridgeBurn, _ := abi.NewBridgeBurner(common.HexToAddress(p.Burner.Address), burnClient)

			etherLocked, _ := etherClient.BalanceAt(ctx, common.HexToAddress(p.Ether.Address), nil)
			tokenAddr, _ := bridgeBurn.Token(nil)
			token, _ := abi.NewWrappedToken(tokenAddr, burnClient)
			tokenMinted, _ := token.TotalSupply(nil)

			fmt.Printf("%s => %s (%s)\n", p.Ether.Client, p.Burner.Client, p.Name)
			valid = valid && etherLocked.Cmp(tokenMinted) == 0
			fmt.Printf("%s\n%s\n%s\n%s\n%t\n\n", p.Ether.Address, etherLocked, p.Burner.Address, tokenMinted, etherLocked.Cmp(tokenMinted) == 0)
		case pairLockBurn:
			lockClient := client[p.Locker.Client]
			bridgeLocker, _ := abi.NewBridgeLocker(common.HexToAddress(p.Locker.Address), lockClient)

			burnClient := client[p.Burner.Client]
			bridgeBurn, _ := abi.NewBridgeBurner(common.HexToAddress(p.Burner.Address), burnClient)

			lockTokenAddr, _ := bridgeLocker.Token(nil)
			lockToken, _ := abi.NewIERC20(lockTokenAddr, lockClient)
			tokenLocked, _ := lockToken.BalanceOf(nil, common.HexToAddress(p.Locker.Address))

			tokenAddr, _ := bridgeBurn.Token(nil)
			token, _ := abi.NewWrappedToken(tokenAddr, burnClient)
			tokenMinted, _ := token.TotalSupply(nil)

			fmt.Printf("%s => %s (%s)\n", p.Locker.Client, p.Burner.Client, p.Name)
			valid = valid && tokenLocked.Cmp(tokenMinted) == 0
			fmt.Printf("%s\n%s\n%s\n%s\n%t\n\n", p.Locker.Address, tokenLocked, p.Burner.Address, tokenMinted, tokenLocked.Cmp(tokenMinted) == 0)
		}
	}

	fmt.Printf("verify result: %t\n", valid)
}
