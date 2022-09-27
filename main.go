package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	todo "github.com/gsagirlar/go-solidty/gen"
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

	//address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	//balance, err := client.BalanceAt(context.Background(), address, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(balance)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//nonce, err := client.PendingNonceAt(context.Background(), address)
	//if err != nil {
	//	log.Fatal(err)
	//}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	/*
		type Todo struct {
			//caller is related to all different function we implemented inside solidity with "view" keyword: list, get
			TodoCaller     // Read-only binding to the contract
			//all other functions: add, update, remove, toggle
			TodoTransactor // Write-only binding to the contract
			//emitters of events
			TodoFilterer   // Log filterer for contract events
		}
	*/

	//this is the deployed contract's address via deploy/main.go
	contractAddress := common.HexToAddress("0x82AC2DE4832f39A563d7d266C0a044e4450c88D7")
	todoStructure, err := todo.NewTodo(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	transOpt, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	transOpt.GasLimit = uint64(3000000)
	transOpt.GasPrice = gasPrice

	//txAdd, err := todoStructure.Add(transOpt, "gokhansa")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Add tx hash \n", txAdd.Hash())

	//txUpdate, err := todoStructure.Update(transOpt, big.NewInt(0), "Gokhan Sagirlar", false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(txUpdate.Hash())
	//
	//txToggle, err := todoStructure.Toggle(transOpt, big.NewInt(1))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Toggle tx: ", txToggle)

	//txDelete, err := todoStructure.Remove(transOpt, big.NewInt(1))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Delete tx: ", txDelete)

	bindCallOption := bind.CallOpts{
		From: crypto.PubkeyToAddress(key.PrivateKey.PublicKey),
		//BlockNumber: nil,
		//Context: context.Background(),
	}
	txList, err := todoStructure.List(&bindCallOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List tasks transaction:", txList)

}
