package chapter1

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp 		int64
	Data 			[]byte
	PreblockHash	[]byte
	Hash 			[]byte
}

func NewBlock(data string,preHash []byte)*Block{
	block := &Block{Timestamp:time.Now().Unix(),Data:[]byte(data),PreblockHash:preHash}
	block.setHash()
	return block
}

func (b *Block) setHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	header := bytes.Join([][]byte{b.PreblockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",[]byte{})
}

