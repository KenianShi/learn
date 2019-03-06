package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"]="A"
	m["b"]="B"
	fmt.Println(len(m))
}
