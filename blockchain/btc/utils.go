package btc

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(data int64 ) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,data)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
