package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	test := flag.NewFlagSet("addblock",flag.ExitOnError)
	fmt.Println(os.Args)
	switch os.Args[2] {
	case "aa":
		fmt.Println("aba")
	default:
		fmt.Println(os.Args[2])
	}
	fmt.Println(test.Parsed())

	fmt.Println(test.Args())

	//fmt.Println(flag.Usage())




}
