package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var goerliURL = "https://goerli.infura.io/v3/52284e357b6a4a8c91ac6962cef491c5"

func main() {
	b, err := os.ReadFile("./wallet/UTC--2022-09-23T14-29-19.095522000Z--5f669bc3d99aab9f84a121693f528517c7b496b3")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(goerliURL)
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

}
