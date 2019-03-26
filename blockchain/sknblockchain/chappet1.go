package main

import (
	"shikenian/learn/blockchain/sknblockchain/chapter1"
	"fmt"
)

func main() {
	bc := chapter1.NewBlockchain()
	bc.AddBlock("Alice send 1 BTC to Bob")
	bc.AddBlock("Bob send 2 BTC to Coco")

	for _,block := range bc.Blocks{
		fmt.Printf("PrevHash:%x.Hash:%x,data:%s. \n",block.PrevBlockHash,block.Hash,block.Data)
	}

}
