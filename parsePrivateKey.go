package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	file     = flag.String("file", "", "file")
	password = flag.String("password", "", "password")
)

func init() {
	flag.Parse()
}

func main() {
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		flag.Usage()
		os.Exit(1)
	}

	keyjson, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	key, err := keystore.DecryptKey(keyjson, *password)
	if err != nil {
		panic(err)
	}

	address := key.Address.Hex()
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	fmt.Printf("Address: %s\nPrivateKey: %s\n",
		address,
		privateKey,
	)
}

//--file /opt/etherAccount/node1/keystore/UTC--2023-07-18T06-10-46.187084661Z--0b587ffd0bba122fb5ddc19ad6eeceb1d2dbbff7 --password yu201219jing
