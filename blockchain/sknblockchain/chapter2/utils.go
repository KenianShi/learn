package chapter2

import (
	"bytes"
	"encoding/binary"
	"log"
)

func checkErr(err error){
	if err != nil {
		log.Panic(err)
	}
}

func IntToHex(num int64)[]byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	checkErr(err)
	return buff.Bytes()
}
