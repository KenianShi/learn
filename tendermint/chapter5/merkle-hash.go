package main

import (
	"fmt"
	"github.com/tendermint/tendermint/crypto/merkle"
)


//之前的tendermint的需要实现Hash方法
//type sh struct {
//	value string
//}
//
//
//func (h sh)Hash() []byte{
//	return tmhash.Sum([]byte(h.value))
//}


//新版的tendermint使用merkle tree，其需要进行merkle tree计算的值都需要序列化成[]byte，
//其中指的是切片里的值以及map里的value
func main() {
	a := make([][]byte,10)			//省略 cap和len的值相等
	a = [][]byte{[]byte("A"),[]byte("B"),[]byte("C")}
	hashByte := merkle.SimpleHashFromByteSlices(a)
	fmt.Printf("%x\n",hashByte)
	a = append(a, []byte("D"))

	//b := map[string]int{"tom":99,"mike":80,"linda":92,"lizzie":66}
	hashByte = merkle.SimpleHashFromByteSlices(a)
	fmt.Printf("%x\n",hashByte)

	m := make(map[string][]byte)
	m["A"]=[]byte("90~100")
	b := merkle.SimpleHashFromMap(m)	//map的value需要转成[]byte
	fmt.Printf("%x\n",b)

}