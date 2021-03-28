---
description: Tutorial on how to load an Ethereum account with Go.
---

# Accounts

Accounts on Ethereum are either wallet addresses or smart contract addresses. They look like `0x8c39f43BDB1a7315aA15b861641d093Bd4F43dD1` and they're what you use for sending ETH to another user and also are used for referring to a smart contract on the blockchain when needing to interact with it. They are unique and are derived from a private key. We'll go more in depth into private/public key pairs in later sections.

In order to use account addresses with go-ethereum, you must first convert them to the go-ethereum `common.Address` type.

```go
address := common.HexToAddress("0x380347b99285a3c7fEE2489A0A6EF9cf018589F1")

fmt.Println(address.Hex()) // 0xe92A52398E068941D9aC03E001e14aF636bcB2F3
```

Pretty much you'd use this type anywhere you'd pass an ethereum address to methods from go-ethereum. Now that you know the basics of accounts and addresses, let's learn how to retrieve the ETH account balance in the next section.

---

### Full code

[address.go](https://github.com/Browser-Coin/ethereum-development-with-go-book/blob/master/code/address.go)

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x380347b99285a3c7fEE2489A0A6EF9cf018589F1")

	fmt.Println(address.Hex())        // 0xe287F9B9C1759903840aC5B139739826535dA471
	fmt.Println(address.Hash().Hex()) // 0x000000000000000000000000000287F9B9C1759903840aC5B139739826535dA471
	fmt.Println(address.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
```
