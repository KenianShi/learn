package main

import (
	"shikenian/learn/blockchain/chapter2"
	"fmt"
)

func main() {
	bc := chapter2.NewBlockchain()
	bc.AddBlock("Bob send a1 BTC to Alice" )
	bc.AddBlock("CC write some code in the blockchain")
	bc.AddBlock("this is the third Block")

	for _,block := range bc.Blocks{
		pow := chapter2.NewProofOfWork(block)
		isValid := pow.Validate()
		fmt.Printf("PreHash: %x \nHash: %x \nData: %s \nnonce: %d \nisValide:%v \n \n \n",block.PreHash,block.Hash,block.Data,block.Nonce,isValid)
	}

}
