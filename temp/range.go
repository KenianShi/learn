package main

import "fmt"

func main() {
	array := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	m := map[string]string{}
	m["a"] = "A"
	m["b"] = "B"
	m["c"] = "C"
	m["d"] = "D"
	for a, b := range array {
		fmt.Println(a, b)
	}

}
