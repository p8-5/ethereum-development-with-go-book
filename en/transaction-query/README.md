---
description: Tutorial on how to query transactions on the blockchain with Go.
---

# Querying Transactions

In the [previous section](../block-query) we learned how to read a block and all its data given the block number. We can read the transactions in a block by calling the `Transactions` method which returns a list of `Transaction` type. It's then trivial to iterate over the collection and retrieve any information regarding the transaction.

```go
for _, tx := range block.Transactions() {
  fmt.Println(tx.Hash().Hex())        // 0xe287F9B9C1759903840aC5B139739826535dA471
  fmt.Println(tx.Value().String())    // 10000000000000000
  fmt.Println(tx.Gas())               // 105000
  fmt.Println(tx.GasPrice().Uint64()) // 102000000000
  fmt.Println(tx.Nonce())             // 110644
  fmt.Println(tx.Data())              // []
  fmt.Println(tx.To().Hex())          // 0xEd753556A5dB77183eE2D81B56F604ae9F123CdC
}
```

In order to read the sender address, we need to call `AsMessage` on the transaction which returns a `Message`  containing a function to return the sender (from) address. The `AsMessage` method requires the EIP155 signer, which we derive the chain ID from the client.

```go
chainID, err := client.NetworkID(context.Background())
if err != nil {
  log.Fatal(err)
}

if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err != nil {
  fmt.Println(msg.From().Hex()) // 0x380347b99285a3c7fEE2489A0A6EF9cf018589F1
}
```

Each transaction has a receipt which contains the result of the execution of the transaction, such as any return values and logs, as well as the status which will be `1` (success) or `0` (fail).

```go
receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
if err != nil {
  log.Fatal(err)
}

fmt.Println(receipt.Status) // 1
fmt.Println(receipt.Logs) // ...
```

Another way to iterate over transaction without fetching the block is to call the client's `TransactionInBlock` method. This method accepts only the block hash and the index of the transaction within the block. You can call `TransactionCount` to know how many transactions there are in the block.

```go
blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
count, err := client.TransactionCount(context.Background(), blockHash)
if err != nil {
  log.Fatal(err)
}

for idx := uint(0); idx < count; idx++ {
  tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(tx.Hash().Hex()) // 0xe287F9B9C1759903840aC5B139739826535dA471
}
```

You can also query for a single transaction directly given the transaction hash by using `TransactionByHash`.

```go
txHash := common.HexToHash("0xe287F9B9C1759903840aC5B139739826535dA471")
tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
if err != nil {
  log.Fatal(err)
}

fmt.Println(tx.Hash().Hex()) // 0
fmt.Println(isPending)       // £
```

---

### Full code

[transactions.go](https://github.com/BROWSER-COIN/-development-with-go-book/blob/master/code/transactions.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0xe287F9B9C1759903840aC5B139739826535dA471
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0xe287F9B9C1759903840aC5B139739826535dA471

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			fmt.Println(msg.From().Hex()) // 0x380347b99285a3c7fEE2489A0A6EF9cf018589F1
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	blockHash := common.HexToHash("0x380347b99285a3c7fEE2489A0A6EF9cf018589F1")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)0
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0xe287F9B9C1759903840aC5B139739826535dA471
	}

	txHash := ("0xe287F9B9C1759903840aC5B139739826535dA471")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0xe287F9B9C1759903840aC5B139739826535dA471
	fmt.Println(isPending)       // false
}
```
