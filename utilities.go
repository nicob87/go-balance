package balance

import "github.com/ethereum/go-ethereum/common"

// DeploymentByChainID returns the address of a known balance checker contract, or ErrNoDeployment if no deployment is
// known for chainID
func DeploymentByChainID(chainID uint64) (common.Address, error) {
	if addr, ok := balanceCheckerDeployments[chainID]; ok {
		return addr, nil
	}
	return common.Address{}, ErrNoDeployment
}
