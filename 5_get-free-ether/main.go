package main

import (
	"fmt"
	"log"
	"context"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var  (
	url = "https://goerli.infura.io/v3/f67125134e064cf094e2495c49323c68"
)

func main() {
	// fmt.Println("je ajab")
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN , keystore.StandardScryptP)

	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 	_, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }


	//"34f4bb72ec332be6b8e5e8b09c1b4b94406f8042"
	//"79e4120aff999ec2682ea209f516288ed09a9e76"

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

	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 1:", b2)

}