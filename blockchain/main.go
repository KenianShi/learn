package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("send 2 Btc to Bob")

	for _,block := range bc.blocks{
		fmt.Printf("PreHash:%x,\nhash:%x:\ndata:%s \n",block.PreHash,block.Hash,block.Data)
	}
}
