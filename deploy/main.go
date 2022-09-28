package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	todo "github.com/gsagirlar/go-solidty/gen"
	"log"
	"math/big"
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

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	client, err := ethclient.Dial(goerliURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	//contractadd, tx, _, err := todo.DeployTodo(auth, client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	contractadd, tx, _, err := todo.DeployTodoFactory(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address: ", contractadd.Hex())
	fmt.Println("Transaction Hash: ", tx.Hash().Hex())

}
