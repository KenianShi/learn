package chapter4

import (
	"log"
	"bytes"
	"encoding/binary"
)

func checkErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func intToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	checkErr(err)
	return buff.Bytes()
}
