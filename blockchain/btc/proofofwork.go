package btc

import (
	"math"
	"math/big"
	"bytes"
	"fmt"
	"crypto/sha256"
)

var maxNonce = math.MaxInt64
const targetBit = 24

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBit))				//左移uint(256-targetBit)，减掉targetBit意思是前面空出多少位
														//target值越大，前面空出的数位越多，也就是值越小，难度越高。
	pow := &ProofOfWork{block:block,target:target}
	return pow
}

func (pow ProofOfWork) prepareData(nonce int) []byte{
	byte := bytes.Join([][]byte{
		pow.block.PreHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(targetBit),
		IntToHex(int64(nonce)),
	},[]byte{})
	return byte
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte

	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1{
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

func (pow *ProofOfWork) Validate() bool{
	var hashInt 	big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
