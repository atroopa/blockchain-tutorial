package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {

	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN , keystore.StandardScryptP)

	password := "password"
	a, err := key.NewAccount(password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Address)

}
