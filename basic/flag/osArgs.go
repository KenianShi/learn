package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("the length of os.Args:")
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)

}