package chapter4

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

func NewBlock(data string,prehash []byte) *Block{
	block := &Block{Timestamp:time.Now().Unix(),Prehash:prehash,Data:[]byte(data)}
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func NewGenesisBlock() *Block{
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

func DeserializeBlock(b []byte)*Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}



