package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"log"
	"math/big"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://goerli.infura.io/v3/f67125134e064cf094e2495c49323c68"
)

func main() {
	client, err := ethclient.Dial(url)

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
	fmt.Println(b1)
	fmt.Println("Balance 2:", ConvertToETH(b2), " ETH")

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}

	var amount *big.Int = big.NewInt(110000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//types.NewTx
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	keystoreFilePath := "./wallet/UTC--2023-08-17T15-57-16.336405000Z--6a8ae59389f338d7b87c3f79058405307f7d2589"
	b, err := ioutil.ReadFile(keystoreFilePath)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx send : %s \n \n", tx.Hash().Hex())
}

// Must Import math/big
func ConvertToETH(input *big.Int) *big.Float {
	value := new(big.Float).Quo(new(big.Float).SetInt(input), big.NewFloat(1e18))
	return value
}
