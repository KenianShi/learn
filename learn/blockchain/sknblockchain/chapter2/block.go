package chapter2

import "time"

type Block struct {
	Timestamp 			int64
	Data 				[]byte
	PrevHash 			[]byte
	Hash 				[]byte
	Nonce 				int
}

func NewBlock(data string,prevBlockHash []byte)*Block{
	block := &Block{Timestamp:time.Now().Unix(),Data:[]byte(data),PrevHash:prevBlockHash,}
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Hash= hash
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block{
	return NewBlock("Genesis block",[]byte{})
}
