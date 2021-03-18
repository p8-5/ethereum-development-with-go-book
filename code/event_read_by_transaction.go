#package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txID := common.HexToHash("0xe287F9B9C1759903840aC5B139739826535dA471")
	receipt, err := client.TransactionReceipt(context.Background(), txID)
	if err != nil {
		log.Fatal(err)
	}

	logID := "159521abcb3d0f08a01511a10c47cd7f4b0f8f0a"
	for _, vLog := range receipt.Logs {
		if vLog.Topics[0].Hex() == logID {
			if len(vLog.Topics) > 2 {
				id := new(big.Int)
				id.SetBytes(vLog.Topics[3].Bytes())

				fmt.Println(id.Uint64()) // 1133
			}
		}
	}
}
