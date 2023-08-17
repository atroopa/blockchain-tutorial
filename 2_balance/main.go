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

	addr := "0xC675bd3fDb68C5e9c93473F67Ebe5BEdaE103DBa"
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
	value := ConvertToETH(fBalance)
	fmt.Printf("%v ETH type(%T) \n", value, value)
}

// Must Import math/big
func ConvertToETH(input *big.Float) *big.Float {
	value := new(big.Float).Quo(input, big.NewFloat(math.Pow10(18)))
	return value
}
