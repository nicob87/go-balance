package balance

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// ErrNoDeployment is an error returned when no deployment is known for a specific ethereum chain ID
var ErrNoDeployment = errors.New("no balance checker deployment for chain ID")

// EtherAddress is an address that is used to represent ether in the balance checker contract
var EtherAddress = common.Address{0}

// maps chain ID to known deployed balance checkers
var balanceCheckerDeployments = map[uint64]common.Address{
	// mainnet deployment:
	// https://etherscan.io/address/0x46b5d094bda9714abf2d90dd692ece546b00d9c5#code
	1: common.HexToAddress("0x46b5D094bDa9714abF2d90Dd692EcE546b00d9C5"),
	// https://polygonscan.com/address/0x0356580d92a571302262a2a5cc0b5b6e09e33722#code
	137: common.HexToAddress("0x0356580d92a571302262a2a5cc0b5b6e09e33722"),
}
