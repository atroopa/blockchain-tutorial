package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	fmt.Println("============================================")
	pvk, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(pvk)

	fmt.Println("Generate Private : ")
	fmt.Println(hexutil.Encode(pData))

	fmt.Println("============================================")
	fmt.Println("Generate Public key: ")
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	fmt.Println("============================================")
	fmt.Println("Generate Ethereum Wallet Address: ")
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
