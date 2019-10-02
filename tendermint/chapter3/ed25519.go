package main

import (
	"fmt"
	kf "github.com/tendermint/tendermint/crypto/ed25519"

)

func main() {
	priv := kf.GenPrivKey()
	fmt.Printf("private key: %v \n",priv)

	pub := priv.PubKey()
	fmt.Printf("pubkey: %v \n",pub)

	addr1 := pub.Address()
	fmt.Printf("addr1: %v \n",addr1)
}
