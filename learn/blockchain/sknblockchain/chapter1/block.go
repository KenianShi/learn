package chapter1

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp 			int64
	Data 				[]byte
	PrevBlockHash 		[]byte
	Hash 				[]byte
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	header := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

func NewBlock(data string,prevBlockHash []byte)*Block{
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}



