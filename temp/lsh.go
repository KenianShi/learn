package main

import (
	"math/big"
	"fmt"
)

func main() {

	target := big.NewInt(1)
	target.Lsh(target, 4)
	fmt.Println(target)
}