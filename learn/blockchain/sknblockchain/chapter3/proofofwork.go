package chapter3

import (
	"math/big"
	"math"
	"fmt"
	"bytes"
	"crypto/sha256"
)

var maxNonce = math.MaxInt64
const targetBits = 24

type ProofOfWork struct {
	Block			*Block
	Target 			*big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork{
	//return &ProofOfWork{Block:b,Target:targetBits}
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	pow := &ProofOfWork{Block:b,Target:target}
	return pow
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\" \n",pow.Block.Data)
	for nonce < maxNonce{
		data := pow.prapareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target) == -1{
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

func (pow *ProofOfWork) prapareData(nonce int)[]byte{
	data := bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},[]byte{})
	return data
}

func (pow *ProofOfWork) Validate()bool{
	var hashInt big.Int
	data := pow.prapareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.Target) == -1
	return isValid
}