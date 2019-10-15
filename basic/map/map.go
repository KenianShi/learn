package main

import "fmt"

func main() {
	var m = make(map[int]string)
	m[1] = "A"
	m[2] = "B"
	fmt.Println(m)

}
