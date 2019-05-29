package chapter4

import "time"

type Block struct {
	Timestamp	 	int64
	PrevBlockHash 	[]byte
	Hash 			[]byte
	Data 			[]byte
}

func  NewBlock(data,prevBlockHash []byte){
	block := &Block{Timestamp:time.Now().Unix(),PrevBlockHash:prevBlockHash,Data:data}

}

