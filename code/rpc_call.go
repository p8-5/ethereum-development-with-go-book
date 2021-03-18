package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// https://stackoverflow.com/questions/53237759/how-to-correctly-send-rpc-call-using-golang-to-get-smart-contract-owner/53260846#53260846

func main() {
	client, err := rpc.DialHTTP("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	type request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	var result string

	req := request{"0xe92A52398E068941D9aC03E001e14aF636bcB2F3", "0x8da5cb5b"}
	if err := client.Call(&result, "eth_call", req, "latest"); err != nil {
		log.Fatal(err)
	}

	owner := common.HexToAddress(result)
	fmt.Printf("%s\n", owner.Hex()) // 0x2D170ce1F719476FeC1a92856cf632aE93444b41
}
