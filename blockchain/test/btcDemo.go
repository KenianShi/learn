package main

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"time"
	"fmt"
)

type Block struct {
	TimeStamp 			int64
	PreHash 		[]byte
	Hash 			[]byte
	Data 			[]byte
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) SetHash(){
	timeStamp :=[]byte(strconv.FormatInt(b.TimeStamp,10))
	header := bytes.Join([][]byte{timeStamp,b.PreHash,b.Data},[]byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

func NewBlock(data string,preHash []byte)*Block{
	block := &Block{TimeStamp:time.Now().Unix(),PreHash:preHash,Data:[]byte(data),}
	block.SetHash()
	return block
}

func NewGenesisBlock()*Block{
	data := "Genesis Block"
	return NewBlock(data,nil)
}

func (blockchain *Blockchain) AddBlock(data string){
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data,preBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockchain() *Blockchain{
	return &Blockchain{Blocks:[]*Block{NewGenesisBlock()}}
}

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Send 1 BTC to Alice")
	blockchain.AddBlock("Send 2 BTC to Bob")

	for _,block := range blockchain.Blocks{
		fmt.Printf("preHash: %x,\nHash: %x\ndata:%s \n",block.PreHash,block.Hash,block.Data)
	}
}