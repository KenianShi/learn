package chapter3

import (
	"time"
	"bytes"
	"encoding/gob"
)

type Block struct {
	Timestamp 			int64
	PrevBlockHash		[]byte
	Hash 				[]byte
	Data 				[]byte
	Nonce 				int
}

func NewBlock(data string,prevBlockHash []byte) *Block{
	block := &Block{Timestamp:time.Now().Unix(),PrevBlockHash:prevBlockHash,Data:[]byte(data)}
	pow :=NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",nil)
}

func (block *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	checkErr(err)
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	checkErr(err)
	return &block
}