package chapter2

import (
	"math/big"
	"bytes"
	"fmt"
	"math"
	"crypto/sha256"
)

const targetBits = 24
const maxNonce=math.MaxInt64

type ProofOfWork struct {
	block 		*Block
	target 		*big.Int
}

func NewProofOfWork(b *Block)*ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))			//将target向左移动256-target位，也就是256前面有targetBits个0，targetBits越大，则前面的0越多，难度越高
	pow := &ProofOfWork{b,target}
	return pow
}

func (pow *ProofOfWork)prepareData(nonce int)[]byte{
	data := bytes.Join([][]byte{
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(nonce)),
		IntToHex(int64(targetBits)),
		pow.block.PreHash,
		pow.block.Data,

	},[]byte{})
	return data
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\" ...... \n",pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) ==-1 {
			break
		}else {
			nonce++
		}
	}
	fmt.Println()
	fmt.Println()
	return nonce,hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) ==-1
	return isValid

}
