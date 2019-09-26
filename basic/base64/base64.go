package main

import (
	"encoding/base64"
	"fmt"
)

var str = "Y3JlYXRvcg==+"

func main() {
	base64Decode(str)
}


func base64Encode(s string)string{
	byte := []byte(s)
	encoded := base64.StdEncoding.EncodeToString(byte)
	fmt.Println(encoded)
	return encoded
}

func base64Decode(s string) string{
	decodeByte,err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	decode := string(decodeByte)
	fmt.Println(decode)
	return decode
}
