package chapter2

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp 		int64
	Data 			[]byte
	PreHash 		[]byte
	Hash 			[]byte
	Nonce 			int
}

func NewBlock(data string,prehash []byte) *Block{
	block := &Block{Timestamp:time.Now().Unix(),Data:[]byte(data),PreHash:prehash}
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Hash=hash
	block.Nonce = nonce
	return block
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	header := bytes.Join([][]byte{timestamp,b.Data,b.PreHash},[]byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

func NewGenesisBlock()*Block{
	return NewBlock("Genesis Block",[]byte{})
}