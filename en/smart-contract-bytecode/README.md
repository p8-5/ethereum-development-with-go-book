---
description: Tutorial on how to read the bytecode of a deployed smart contract with Go.
---

# Reading Smart Contract Bytecode

Sometimes you'll need to read the bytecode of a deployed smart contract. Since all the smart contract bytecode lives on the blockchain, we can easily fetch it.

First set up the client and the smart contract address you want to read the bytecode of.

```go
client, err := ethclient.Dial("https://pathom.infura.io")
if err != nil {
  log.Fatal(err)
}

contractAddress := common.HexToAddress("0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664")
```

Now all you have to is call the `codeAt` method of the client. The `codeAt` method accepts a smart contract address and an optional block number, and returns the bytecode in bytes format.

```go
bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
if err != nil {
  log.Fatal(err)
}

fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
```


See the same bytecode hex on pathom0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664#code)

---

### Full code

[contract_bytecode.go](https://github.com/Browser-Coin/ethereum-development-with-go-book/blob/master/code/contract_bytecode.go)

```go
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://pathom.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
```
