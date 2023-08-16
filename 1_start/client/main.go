package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var unfuraURL = "https://mainnet.infura.io/v3/f67125134e064cf094e2495c49323c68"
var ganacheURL = "http://127.0.0.1:7545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURL)

	if err != nil {
		log.Fatalf("Error: error to create a ether client -> %v", err)
		return
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("Error: error to get Block -> %v", err)
		return
	}

	fmt.Println(block.Number())
}
