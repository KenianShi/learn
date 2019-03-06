package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is an error")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			fmt.Println("Unknown err")
		} else {
			fmt.Printf("pathErr.Op:%s, \npathErr.Path:%s, \npathErr.Err:%s \n ", pathErr.Op, pathErr.Path, pathErr.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("aaa.txt")

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
