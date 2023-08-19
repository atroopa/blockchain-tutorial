package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

func main() {

	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN , keystore.StandardScryptP)

	password := "password"
	// a, err := key.NewAccount(password)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(a.Address)

	b, err := ioutil.ReadFile("wallet/UTC--2023-08-17T14-53-48.101319400Z--e1a8d0489dbde7b714febf6ad01d5b4db4f54201")

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("private key : \n 		", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("public key : \n 		", hexutil.Encode(pData))

	fmt.Println("address: \n 		", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
