package main

import (
	"fmt"
	kf  "github.com/tendermint/tendermint/crypto/secp256k1"
)

func main() {
	priv := kf.GenPrivKey()
	fmt.Printf("private key ==> %v \n",priv)

	pub := priv.PubKey()
	fmt.Printf("pubkey ==> %v \n",pub)

	addr1 := pub.Address()
	fmt.Printf("address1: %v \n",addr1)

	addr2 := pub.Address()
	fmt.Printf("address2: %v \n",addr2)

}
