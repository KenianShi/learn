package main

import (
	"fmt"
	"strings"
)

const charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

func main() {

	bech := "cosmos1cde9j30h8teunj2fcz422y8lr0e2xu5qxsnd8k"
	one := strings.LastIndexByte(bech, '1')
	data := bech[one+1:]
	hrp := bech[:one]
	decoded, err := toBytes(data)
	if err != nil {
		panic(err)
	}


	checksum := bech[len(bech)-6:]
	fmt.Printf("hrp:%v,decoded:%s\n",hrp,"fdssd")
	fmt.Println("********************************")
	fmt.Println(decoded[:len(decoded)-6])
	a := bech32Checksum(hrp,decoded[:len(decoded)-6])
	fmt.Println("================================")
	fmt.Println(a)
	expected, err := toChars(a)
	if err == nil {
		moreInfo := fmt.Sprintf("Expected %v, got %v.",
			expected, checksum)
		fmt.Println(moreInfo)
	}


}



func toChars(data []byte) (string, error) {
	result := make([]byte, 0, len(data))
	for _, b := range data {
		if int(b) >= len(charset) {
			return "", fmt.Errorf("invalid data byte: %v", b)
		}
		result = append(result, charset[b])
	}
	return string(result), nil
}



func bech32Checksum(hrp string, data []byte) []byte {
	// Convert the bytes to list of integers, as this is needed for the
	// checksum calculation.
	integers := make([]int, len(data))
	for i, b := range data {
		integers[i] = int(b)
	}
	values := append(bech32HrpExpand(hrp), integers...)
	values = append(values, []int{0, 0, 0, 0, 0, 0}...)
	polymod := bech32Polymod(values) ^ 1
	var res []byte
	for i := 0; i < 6; i++ {
		res = append(res, byte((polymod>>uint(5*(5-i)))&31))
	}
	return res
}

func bech32HrpExpand(hrp string) []int {
	v := make([]int, 0, len(hrp)*2+1)
	for i := 0; i < len(hrp); i++ {
		v = append(v, int(hrp[i]>>5))
	}
	v = append(v, 0)
	for i := 0; i < len(hrp); i++ {
		v = append(v, int(hrp[i]&31))
	}
	return v
}

func bech32Polymod(values []int) int {
	chk := 1
	for _, v := range values {
		b := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ v
		for i := 0; i < 5; i++ {
			if (b>>uint(i))&1 == 1 {
				chk ^= gen[i]
			}
		}
	}
	return chk
}

var gen = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}


func toBytes(chars string) ([]byte, error) {
	decoded := make([]byte, 0, len(chars))
	for i := 0; i < len(chars); i++ {
		index := strings.IndexByte(charset, chars[i])
		if index < 0 {
			return nil, fmt.Errorf("invalid character not part of "+
				"charset: %v", chars[i])
		}
		decoded = append(decoded, byte(index))
	}
	return decoded, nil
}