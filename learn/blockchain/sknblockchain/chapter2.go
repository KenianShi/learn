package main

import (
	"shikenian/learn/blockchain/sknblockchain/chapter2"
	"fmt"
)

func main() {
	bc := chapter2.NewBlockchain()

	bc.AddBlock("Alice send 2 BTC to Bob")
	bc.AddBlock("Bob send nothing to Alice")

	for _,block := range bc.Blocks{
		pow := chapter2.NewProofOfWork(block)
		if pow.Validator() {
			fmt.Printf("验证通过，data:%s |||| preHash:%x,  ||||  hash:%x \n",block.Data,block.PrevHash,block.Hash)
		}else{
			fmt.Println("validation failed!")
			break
		}
	}

}