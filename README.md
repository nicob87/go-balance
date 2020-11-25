# BalanceChecker (`go-balance`)

## Contract

The [`BalanceChecker`](https://etherscan.io/address/0x46b5d094bda9714abf2d90dd692ece546b00d9c5#code) contract allows the balances of multiple tokens for multiple addresses to be read with a single `eth_call` instead of several.

Within practical limits on the number of tokens and number of users specified, this can be significantly more effective than multiple individual `balanceOf` calls.

This contract includes a modification that allows ether balance to be read as well, by specifying the token address at a given position in the tokens array as the null address (`address(0)`).

```solidity
interface BalanceChecker {
    function balances(
        address[] calldata users,
        address[] calldata tokens
    )
        external
        view
        returns (uint256[] memory)
}
```

## Library

Contained is the go package `go-balance` which includes go bindings and a simple abstraction around the `BalanceChecker` contract.

Install the package with:

```
go get -u github.com/hrharder/go-balance
```

For usage, see [the included example.](./example/main.go)

## Deployments

Known deployments of BalanceCheckers are listed below.

| ChainID | Address | Date |
| - | - | - |
| `1` | [`0x46b5D094bDa9714abF2d90Dd692EcE546b00d9C5`](https://etherscan.io/address/0x46b5d094bda9714abf2d90dd692ece546b00d9c5) | 2020-11-25 |

## Notes

- Credit to Steve Marx (@smarx) for the Solidity.
- Re-generate the contract bindings as needed with `go generate ./...` (from the package root).