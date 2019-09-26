package chapter2

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(n int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,n)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
