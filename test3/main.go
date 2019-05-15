package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jbenet/go-base58"
	"github.com/tyler-smith/go-bip32"
	"log"
)

// 時間切れで未完成
func main() {
	// bip32
	seed, err := bip32.NewSeed()
	if err != nil {
		log.Fatal(err)
	}

	master, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}

	// m/44'
	key, err := master.NewChildKey(2147483648 + 44)
	if err != nil {
		log.Fatal(err)
	}

	decoded := base58.Decode(key.B58Serialize())
	privateKey := decoded[46:78]
	fmt.Println(hexutil.Encode(privateKey))

	// Hex private key to ECDSA private key
	privateKeyECDSA, err := crypto.ToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// ECDSA private key to hex private key
	privateKey = crypto.FromECDSA(privateKeyECDSA)
	fmt.Println(hexutil.Encode(privateKey))
}


