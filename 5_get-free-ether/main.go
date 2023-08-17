package main

import (
	"fmt"
	"log"
	"math/big"
	"context"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var  (
	url = "https://goerli.infura.io/v3/f67125134e064cf094e2495c49323c68"
)

func main() {
	client , err := ethclient.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	a1 := common.HexToAddress("34f4bb72ec332be6b8e5e8b09c1b4b94406f8042")
	a2 := common.HexToAddress("79e4120aff999ec2682ea209f516288ed09a9e76")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1:", ConvertToETH(b1), " ETH")
	fmt.Println("Balance 2:", ConvertToETH(b2), " ETH")
}

// Must Import math/big
func ConvertToETH(input *big.Int) *big.Float {
	value := new(big.Float).Quo(new(big.Float).SetInt(input), big.NewFloat(1e18))
	return value
}
