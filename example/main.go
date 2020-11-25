package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hrharder/go-balance"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// config (replace if needed)
const ethURI = "http://localhost:8545"
const chainID = 1

// set some users (owner addresses)
var users = []common.Address{
	common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"), // put your address(es) here
}

// set some tokens (token addresses)
var tokens = []common.Address{
	balance.EtherAddress,
	common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), // WETH
	common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), // DAI
	common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"), // USDC
}

func main() {
	eth, err := ethclient.Dial(ethURI)
	must("dial ETH network", err)

	balanceChecker, err := balance.NewChecker(eth, 1)
	must("create balance fetcher", err)

	balances, err := balanceChecker.GetBalances(nil, users, tokens)
	must("fetch balances", err)

	balancesJSON, _ := json.Marshal(balances)
	fmt.Println(string(balancesJSON))
}

func must(s string, err error) {
	if err != nil {
		log.Fatalf("failed to %s: %s", s, err.Error())
	}
}
