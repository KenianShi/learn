package btc

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"

	"encoding/gob"
)

type Block struct {
	Timestamp 		int64
	PreHash 		[]byte
	Hash 			[]byte
	Data 			[]byte
	Nonce			int
}



func (b *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return &block

}


func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	header := bytes.Join([][]byte{b.PreHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(header)
	b.Hash=hash[:]
}

func NewBlock(data string,prehash []byte)*Block{
	block := &Block{Timestamp:time.Now().Unix(),PreHash:prehash,Data:[]byte(data),}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	return block
}

func NewGenesisBlock()*Block{
	return NewBlock("Genesis block",nil)

}






