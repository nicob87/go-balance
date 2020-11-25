package balance

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Checker is an abstraction over a BalanceChecker contract
type Checker struct {
	bcc *BalanceCheckerCaller
}

// NewChecker returns a new multi-balan
func NewChecker(caller bind.ContractCaller, chainID uint64) (*Checker, error) {
	addr, err := DeploymentByChainID(chainID)
	if err != nil {
		return nil, err
	}
	bcc, err := NewBalanceCheckerCaller(addr, caller)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to balance checker: %w", err)
	}
	return &Checker{bcc}, nil
}

// GetBalances gets the balance of each of tokens for each of users
func (c *Checker) GetBalances(opts *bind.CallOpts, users, tokens []common.Address) (MultiBalances, error) {
	result, err := c.bcc.Balances(opts, users, tokens)
	if err != nil {
		return nil, err
	}
	if (len(users))*(len(tokens)) != len(result) {
		return nil, fmt.Errorf("bad results length: expected (%v) != actual (%v)", (len(users))*(len(tokens)), len(result))
	}

	balances := make(MultiBalances)
	for i := 0; i < len(users); i++ {
		balances[users[i]] = make(BalancesByToken)
		for j := 0; j < len(tokens); j++ {
			balances[users[i]][tokens[j]] = result[j+len(tokens)*i]
		}
	}
	return balances, nil
}
