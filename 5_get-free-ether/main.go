package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	url = "http://127.0.0.1:7545" // آدرس و پورت Ganache
)

func main() {
	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	a1 := common.HexToAddress("0x6137BA6b5003961d98d866EBAaCdC54E549E7601") // آدرس اکانت اول در Ganache
	a2 := common.HexToAddress("0x1aa2F228C93Af553dC81a8D068FAe1eD2F04eA95") // آدرس اکانت دوم در Ganache

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

	//keystoreFilePath := "./wallet/UTC--2023-08-17T15-56-10.974441400Z--34f4bb72ec332be6b8e5e8b09c1b4b94406f8042"
	//b, err := ioutil.ReadFile("./wallet/UTC--2023-08-17T15-56-10.974441400Z--34f4bb72ec332be6b8e5e8b09c1b4b94406f8042")
	//if err != nil {
	//	log.Fatal(err)
	//}

	priv := "0xe5e893e2cafb0ce25d57e83501dea1f017f9571f4c37b8337db4c2662775d742"

	key, err := crypto.HexToECDSA(priv)

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), key)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx send : %s", tx.Hash().Hex())
}

// ConvertToETH Must Import math/big
func ConvertToETH(input *big.Int) *big.Float {
	value := new(big.Float).Quo(new(big.Float).SetInt(input), big.NewFloat(1e18))
	return value
}
