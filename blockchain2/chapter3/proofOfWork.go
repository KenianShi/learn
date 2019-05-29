package chapter3

import (
	"math/big"
	"math"
	"fmt"
	"bytes"
	"crypto/sha256"
)

const targetBits  = 24
var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block 	*Block
	target 	*big.Int
}

func NewProofofWork(block *Block)*ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	pow := &ProofOfWork{block,target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int)[]byte{
	data := bytes.Join(
		[][]byte{
		pow.block.Prehash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(nonce)),
		IntToHex(int64(targetBits)),
	},[]byte{})
	return data
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\" \n ",pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1{
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}