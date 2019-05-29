package chapter3

import (
	"math"
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

const targetBits = 24
//const targetBits = 4


const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block		*Block
	target 		*big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	pow := &ProofOfWork{b,target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join([][]byte{
		IntToHex(int64(pow.block.Timestamp)),
		pow.block.Prehash,
		pow.block.Data,
		IntToHex(int64(nonce)),
		IntToHex(targetBits),
	},[]byte{})
	return data
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining block containing \"%s\" \n :",pow.block.Data)
	for nonce < maxNonce{
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x ",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		}else{
			nonce++
		}
	}
	fmt.Printf("\n \n")
	return nonce,hash[:]
}

func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) ==-1
}