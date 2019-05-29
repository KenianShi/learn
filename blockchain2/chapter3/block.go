package chapter3

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp 		int64
	Prehash 		[]byte
	Data 			[]byte
	Hash 			[]byte
	Nonce 			int
}

func NewBlock(data string,preBlockHash []byte)*Block {
	block := &Block{Timestamp: time.Now().Unix(), Prehash: preBlockHash, Data: []byte(data)}
	pow := NewProofofWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesis() *Block{
	return NewBlock("Genesis Block",[]byte{})
}

func (b *Block) Serialize()[]byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte)*Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}