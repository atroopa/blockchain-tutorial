package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
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

	fmt.Println("the Block Number is : ", block.Number())

	addr := "0x728Dc029218C75AAcfB21DC663972745a85d00c1"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		log.Fatalf("Error: error to get Balance -> %v", err)
		return
	}

	fmt.Println("the Balance: ", balance)
	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Printf("big number : %v   type(%T) \n", fBalance, fBalance)
	// 10*10*10*10*10*...18
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("%v ETH type(%T) \n", value, value)
}

// func ConverToETH(input float64){
// 	value := new(big.Float).Quo(input, big.NewFloat(math.Pow10(18)))
// 	return value
// }
