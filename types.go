package balance

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MultiBalances is a map of an account (user) to balances by token address
type MultiBalances map[common.Address]BalancesByToken

// BalancesByToken is a map of token address to a unit value for the balance
type BalancesByToken map[common.Address]*big.Int
