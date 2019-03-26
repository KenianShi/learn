package chapter4

import (
	"time"
	"crypto/sha256"
	"bytes"
	"encoding/gob"
)

type Block struct {
	Timestamp 			int64
	Transactions 		[]*Transaction
	PreBlockHash 		[]byte
	Hash 				[]byte
	Nonce 				int
}

func (b *Block) HashTransaction() []byte{
	var txHashes [][]byte
	var txHash [32]byte

	for _,tx := range b.Transactions{
		txHashes = append(txHashes,tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))
	return txHash[:]
}

func NewBlock(transactions []*Transaction,preBlockHash []byte) *Block{
	block := &Block{Timestamp:time.Now().Unix(),Transactions:transactions,PreBlockHash:preBlockHash,Hash:[]byte{},Nonce:0}
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func NewGenesisBlock(coinbase *Transaction)*Block{
	return NewBlock([]*Transaction{coinbase},[]byte{})
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


