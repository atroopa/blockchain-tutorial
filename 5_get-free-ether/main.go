package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcURL = "http://127.0.0.1:7545" // URL of your Ethereum node (e.g., Ganache)
)

func main() {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Sender's private key
	senderPrivKeyHex := "0x8ba0df86dee03cf98c77aba636fdbc5e3c3993cc2173d8dbc3d38f20144c8833"
	senderPrivKey, err := crypto.HexToECDSA(strings.TrimPrefix(senderPrivKeyHex, "0x"))
	if err != nil {
		log.Fatalf("senderPrivKey : ", err)
	}

	// Receiver's address
	receiverAddr := common.HexToAddress("0x48EDbf3f0d8a7971D330d18b31a19B20D8fd53D5")

	// Set up transactor using sender's private key
	auth := bind.NewKeyedTransactor(senderPrivKey)

	// Get the nonce for the sender's address
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("nonce: ", err)
	}

	// Amount to send (in wei)
	amount := big.NewInt(1000000000000000000) // 1 Ether

	// Gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("gasPrice : ", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, receiverAddr, amount, 21000, gasPrice, nil)

	// Chain ID (1 for Ethereum mainnet, 3 for Ropsten, etc.)
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatalf("chainID: ", err)
	// }
	chainID := big.NewInt(1337) // this is ganache chain id

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderPrivKey)
	if err != nil {
		log.Fatalf("signedTx: ", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("SendTransaction: ", err)
	}

	fmt.Printf("Transaction Hash: 0x%x\n", signedTx.Hash())
}
