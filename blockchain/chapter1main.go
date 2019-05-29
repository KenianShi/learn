package main

import "shikenian/learn/blockchain/chapter1"
import "fmt"

func main() {
	bc := chapter1.NewBlockchain()
	bc.AddBlock("Alice Send 5 BTC to Bob")
	bc.AddBlock("Bob send 2 BTC to Tom")
	bc.AddBlock("tom have 2 BTC")

	for _,block := range bc.Block{
		fmt.Printf("PreHash: %x \n",block.PreblockHash)
		fmt.Printf("Hash: %x \n",block.Hash)
		fmt.Printf("Data: %x \n",block.Data)
		fmt.Println()
	}

}

